// Package main provides ...
package main

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"

	cmd_demo "github.com/deepzz0/appdemo/api/cmd-demo"
	"github.com/deepzz0/appdemo/cmd/demo/rpc"
	"github.com/deepzz0/appdemo/cmd/demo/swag"
	"github.com/deepzz0/appdemo/cmd/demo/user"
	"github.com/deepzz0/appdemo/pkg/config"
	"github.com/deepzz0/appdemo/pkg/i18n"
	"github.com/deepzz0/appdemo/pkg/middleware/language"
	"github.com/deepzz0/appdemo/pkg/middleware/session"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// @title APP Demo API
// @version 1.0
// @description This is a sample server celler server.

// @BasePath /api

func main() {
	fmt.Println("Hi, it's App " + config.Conf.DemoApp.Name)

	endRun := make(chan bool, 1)

	runHTTPServer(endRun)
	runGRPCServer(endRun)
	<-endRun
}

func runHTTPServer(endRun chan bool) {
	if !config.Conf.DemoApp.EnableHTTP {
		return
	}

	if config.Conf.RunMode == config.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.Default()
	// load html
	glob := filepath.Join(config.WorkDir, "website", "*.html")
	e.LoadHTMLGlob(glob)

	// middleware
	e.Use(language.Middleware(language.Options{
		CookieName: "lang",
		Default:    i18n.GetDefaultLang(),
		Supported:  i18n.GetSupportedLang(),
	}))
	e.Use(session.Middleware(session.Options{
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
	address := fmt.Sprintf(":%d", config.Conf.DemoApp.HTTPPort)
	go e.Run(address)
	fmt.Println("HTTP server running on: " + address)
}

func runGRPCServer(endRun chan bool) {
	if !config.Conf.DemoApp.EnableGRPC {
		return
	}

	address := fmt.Sprintf(":%v", config.Conf.DemoApp.GRPCPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	cmd_demo.RegisterUserServer(s, &rpc.UserSrv{})

	go s.Serve(lis)
	fmt.Println("GRPC server running on: " + address)
}
