// Package config provides ...
package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	// Conf config instance
	Conf Config

	// ModeDev run mode as development
	ModeDev = "dev"
	// ModeProd run mode as production
	ModeProd = "prod"
	// Gopath go path
	Gopath = os.Getenv("GOPATH")
)

// Mode run mode
type Mode struct {
	EnableHTTP bool `yaml:"enablehttp"`
	HTTPPort   int  `yaml:"httpport"`
}

// Database sql database
type Database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

// Config app config
type Config struct {
	RunMode  string   `yaml:"runmode"`
	AppName  string   `yaml:"appname"`
	AppMode  Mode     `yaml:"appmode"`
	Database Database `yaml:"database"`
}

// load config file
func init() {
	// compatibility linux and windows
	path := filepath.Join("conf", "app.yml")
	if Gopath != "" {
		path = filepath.Join(Gopath, "src", "github.com",
			"deepzz0", "appdemo", path)
	}

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
}
