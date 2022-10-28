// Package language provides ...
package language

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Options 语言选项
type Options struct {
	CookieName string
	Default    string
	Supported  []string
}

// isExist language
func (opts Options) isExist(l string) bool {
	for _, v := range opts.Supported {
		if v == l {
			return true
		}
	}
	return false
}

// Middleware set language
func Middleware(opts Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang, err := c.Cookie(opts.CookieName)
		// found cookie
		if err == nil {
			c.Set(opts.CookieName, lang)
			return
		}
		// set cookie
		al := strings.ToLower(c.GetHeader("Accept-Language"))
		if al != "" {
			// choose default if not supported
			lang = opts.Default

			langs := strings.Split(al, ",")
			for _, v := range langs {
				if opts.isExist(v) {
					lang = v
					break
				}
			}
		} else {
			lang = opts.Default
		}
		c.SetCookie(opts.CookieName, lang, 86400*365, "/", "", false, false)
		c.Set(opts.CookieName, lang)
	}
}
