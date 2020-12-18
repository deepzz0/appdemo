// Package mid provides ...
package mid

import (
	"strings"

	"github.com/deepzz0/appdemo/pkg/i18n"

	"github.com/gin-gonic/gin"
)

// SetLanguage set language
func SetLanguage(c *gin.Context) {
	lang, err := c.Cookie("lang")
	// found cookie
	if err == nil {
		c.Set("lang", lang)
		return
	}
	// set cookie
	al := strings.ToLower(c.GetHeader("Accept-Language"))
	if al != "" {
		// choose default if not supported
		lang = i18n.DefaultLang

		langs := strings.Split(al, ",")
		for _, v := range langs {
			if i18n.Success.IsExist(v) {
				lang = v
				break
			}
		}
	} else {
		lang = i18n.DefaultLang
	}
	c.SetCookie("lang", lang, 86400*365, "/", "", false, false)
	c.Set("lang", lang)
}
