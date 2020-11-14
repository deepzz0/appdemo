// Package main provides ...
package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/deepzz0/appdemo/pkg/api/swag"
	"github.com/deepzz0/appdemo/pkg/api/user"
	"github.com/deepzz0/appdemo/pkg/config"
	"github.com/deepzz0/appdemo/pkg/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi, it's App Demo")

	if config.Conf.RunMode == config.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.Default()
	// load html
	glob := filepath.Join(config.WorkDir, "website", "*.html")
	e.LoadHTMLGlob(glob)
	// session store
	store := cookie.NewStore([]byte("ZGlzvcmUoMTAsICI="))
	store.Options(sessions.Options{
		MaxAge:   86400 * 30,
		Path:     "/",
		Secure:   config.Conf.RunMode == config.ModeProd,
		HttpOnly: true,
	})
	e.Use(sessions.Sessions("SESSION_ID", store))

	// middleware
	e.Use(middleware.SetLanguage)

	// router
	e.GET("/swagger/*any", swag.HandleSwagger)

	// static files, page
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// api router
	api := e.Group("/api")
	{
		api.POST("/register", user.HandleRegister)
		api.POST("/login", user.HandleLogin)
	}
	authz := api.Use(user.AuthFilter)
	{
		authz.GET("/logout", user.HandleLogout)
		authz.GET("/userinfo", user.HandleUserInfo)
	}

	// start
	if config.Conf.AppDemo.EnableHTTP {
		e.Run(fmt.Sprintf(":%d", config.Conf.AppDemo.HTTPPort))
	}
}
