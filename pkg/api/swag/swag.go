// Package swag provides ...
package swag

import (
	_ "github.com/deepzz0/appdemo/docs" // docs

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// HandleSwagger swagger docs
var HandleSwagger = ginSwagger.WrapHandler(swaggerFiles.Handler)
