// Package mid provides ...
package mid

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionOpts 设置选项
type SessionOpts struct {
	Secure bool
	Secret []byte
	// redis store
	RedisAddr string
	RedisPwd  string
}

// SessionMiddleware session中间件
func SessionMiddleware(opts SessionOpts) gin.HandlerFunc {
	store := cookie.NewStore(opts.Secret)
	store.Options(sessions.Options{
		MaxAge:   86400 * 30,
		Path:     "/",
		Secure:   opts.Secure,
		HttpOnly: true,
	})
	return sessions.Sessions("SESSIONID", store)
}
