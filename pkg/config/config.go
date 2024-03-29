// Package config provides ...
package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	// Conf config instance
	Conf Config

	// ModeDev run mode as development
	ModeDev = "dev"
	// ModeProd run mode as production
	ModeProd = "prod"
	// WorkDir workspace dir
	WorkDir string
)

// Mode run mode
type Mode struct {
	Name       string
	EnableHTTP bool
	HTTPPort   int
	EnableGRPC bool
	GRPCPort   int
}

// Database sql database
type Database struct {
	Driver string
	Source string
}

// CacheRedis cache redis
type CacheRedis struct {
	Host     string
	Password string
	DB       int
}

// Config app config
type Config struct {
	RunMode  string
	AppName  string
	Database Database
	DemoApp  Mode
}

// load config file
func init() {
	// compatibility linux and windows
	var err error
	WorkDir = workDir()
	path := filepath.Join(WorkDir, "conf", "app.yml")

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		panic(err)
	}
	// read run mode from env
	if runmode := os.Getenv("RUN_MODE"); runmode != "" {
		if runmode != ModeDev && runmode != ModeProd {
			panic("invalid RUN_MODE from env: " + runmode)
		}
		Conf.RunMode = runmode
	}
	// read env
	readDBEnv()
}

func readDBEnv() {
	key := strings.ToUpper(Conf.AppName) + "_DB_DRIVER"
	if d := os.Getenv(key); d != "" {
		Conf.Database.Driver = d
	}
	key = strings.ToUpper(Conf.AppName) + "_DB_SOURCE"
	if s := os.Getenv(key); s != "" {
		Conf.Database.Source = s
	}
}
