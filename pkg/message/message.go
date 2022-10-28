// Package message provides ...
package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespCode api response code
type RespCode interface {
	Tr(lang string, args ...interface{}) string
	StatusCode() int
}

// Message response message
type Message struct {
	Code  RespCode    `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// JSON response as json
func (msg *Message) JSON(c *gin.Context) {
	status := msg.Code.StatusCode()
	if status != http.StatusOK && msg.Error == "" {
		lang := c.Keys["lang"].(string)
		msg.Error = msg.Code.Tr(lang)
	}
	c.JSON(status, msg)
}
