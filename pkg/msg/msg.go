// Package msg provides ...
package msg

import (
	"github.com/gin-gonic/gin"
)

// RespCode api response code
type RespCode interface {
	Tr(lang string) string
	StatusCode() int
	IsExist(lang string) bool
}

// Message response message
type Message struct {
	Code  RespCode    `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// JSON response as json
func (msg *Message) JSON(c *gin.Context) {
	if msg.Error == "" {
		lang := c.Keys["lang"].(string)
		msg.Error = msg.Code.Tr(lang)
	}
	c.JSON(msg.Code.StatusCode(), msg)
}
