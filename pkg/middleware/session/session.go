// Package session provides ...
package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Options 设置选项
type Options struct {
	Secure bool
	Secret []byte
	// redis store
	RedisAddr string
	RedisPwd  string
}

// Middleware session中间件
func Middleware(opts Options) gin.HandlerFunc {
	store := cookie.NewStore(opts.Secret)
	store.Options(sessions.Options{
		MaxAge:   86400 * 30,
		Path:     "/",
		Secure:   opts.Secure,
		HttpOnly: true,
	})
	return sessions.Sessions("SESSIONID", store)
}
