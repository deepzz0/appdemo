// Package user provides ...
package user

import (
	"encoding/gob"
	"net/http"
	"strings"

	"github.com/deepzz0/appdemo/pkg/i18n"
	"github.com/deepzz0/appdemo/pkg/message"
	"github.com/deepzz0/appdemo/pkg/model/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	gob.Register(&user.User{})
}

// RegisterRoutes register routes
func RegisterRoutes(group gin.IRoutes) {
	group.POST("/register", handleRegister)
	group.POST("/login", handleLogin)
}

// RegisterRoutesAuthz register routes
func RegisterRoutesAuthz(group gin.IRoutes) {
	group.GET("/logout", handleLogout)
	group.GET("/userinfo", handleUserInfo)
}

// AuthFilter auth filter
func AuthFilter(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Abort()
		c.String(http.StatusUnauthorized, "unauthorized user")
		return
	}
	c.Set("user", user)

	c.Next()
}

type registerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// handleRegister register user
// @Summary register user
// @Description register user
// @Tags User
// @Accept json
// @Produce json
// @Param req body registerReq true "register info"
// @Success 200 {object} message.Message
// @Router /register [post]
func handleRegister(c *gin.Context) {
	msg := &message.Message{}
	defer msg.JSON(c)

	var req registerReq
	if err := c.BindJSON(&req); err != nil {
		msg.Code = i18n.ErrBadRequest
		return
	}
	// check params
	if !strings.Contains(req.Username, "@") {
		msg.Code = i18n.ErrInvalidUsername
		return
	}
	if len(req.Password) < 6 {
		msg.Code = i18n.ErrInvalidPassword
		return
	}
	// is exists
	exist, err := user.IsExistUser(req.Username)
	if err != nil {
		msg.Code = i18n.ErrSystemInternal
		return
	}
	if exist {
		msg.Code = i18n.ErrAlreadyExistUser
		return
	}
	// create user
	u := &user.User{
		Username: req.Username,
		Password: req.Password,

		UserAgent: c.Request.UserAgent(),
		CreatedIP: c.ClientIP(),
	}
	err = user.InsertUser(u)
	if err != nil {
		msg.Code = i18n.ErrSystemInternal
		return
	}
	msg.Code = i18n.Success
}

type loginReq registerReq

// handleLogin login user
// @Summary login user
// @Description login user
// @Tags User
// @Accept json
// @Produce json
// @Param req body loginReq true "login info"
// @Success 200 {object} message.Message
// @Router /login [post]
func handleLogin(c *gin.Context) {
	msg := &message.Message{}
	defer msg.JSON(c)

	var req registerReq
	if err := c.BindJSON(&req); err != nil {
		msg.Code = i18n.ErrBadRequest
		return
	}
	// check params
	if !strings.Contains(req.Username, "@") {
		msg.Code = i18n.ErrInvalidUsername
		return
	}
	if len(req.Password) < 6 {
		msg.Code = i18n.ErrInvalidPassword
		return
	}
	// valid user
	u, err := user.SelectUserByUsername(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg.Code = i18n.ErrNotFoundUser
		} else {
			msg.Code = i18n.ErrSystemInternal
		}
		return
	}
	if u.Password != req.Password {
		msg.Code = i18n.ErrIncorrectPassword
		return
	}
	// login success
	session := sessions.Default(c)
	session.Set("user", u)
	err = session.Save()
	msg.Code = i18n.Success
}

// handleLogout logout user
// @Summary log out
// @Description log out
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} message.Message
// @Router /logout [get]
func handleLogout(c *gin.Context) {
	msg := &message.Message{}
	defer msg.JSON(c)

	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	msg.Code = i18n.Success
}

// handleUserInfo user info
// @Summary user info
// @Description user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} user.User
// @Router /userinfo [get]
func handleUserInfo(c *gin.Context) {
	msg := &message.Message{}
	defer msg.JSON(c)

	u := c.Keys["user"].(*user.User)
	msg.Code = i18n.Success
	msg.Data = u.ForShow()
}
