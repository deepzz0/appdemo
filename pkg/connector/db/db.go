// Package db provides ...
package db

import (
	"github.com/deepzz0/appdemo/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm db
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.Conf.Database.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
