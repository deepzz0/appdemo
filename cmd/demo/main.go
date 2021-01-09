// Package main provides ...
package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/deepzz0/appdemo/pkg/config"
	"github.com/deepzz0/appdemo/pkg/core/demo/swag"
	"github.com/deepzz0/appdemo/pkg/core/demo/user"
	"github.com/deepzz0/appdemo/pkg/i18n"
	"github.com/deepzz0/appdemo/pkg/mid"

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

	// middleware
	e.Use(mid.LangMiddleware(mid.LangOpts{
		CookieName: "lang",
		Default:    i18n.GetDefaultLang(),
		Supported:  i18n.GetSupportedLang(),
	}))
	e.Use(mid.SessionMiddleware(mid.SessionOpts{
		Secure: config.Conf.RunMode == config.ModeProd,
		Secret: []byte("ZGlzvcmUoMTAsICI="),
	}))

	// swag
	swag.RegisterRoutes(e)

	// static files, page
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// api router
	group := e.Group("/api")
	{
		user.RegisterRoutes(group)
	}
	group = e.Group("/api", user.AuthFilter)
	{
		user.RegisterRoutesAuthz(group)
	}

	// start
	if config.Conf.DemoApp.EnableHTTP {
		e.Run(fmt.Sprintf(":%d", config.Conf.DemoApp.HTTPPort))
	}
}
