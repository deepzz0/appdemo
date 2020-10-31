// Package i18n provides ...
package i18n

import (
	"fmt"
	"net/http"
)

// DefaultLang default lang
const DefaultLang = "zh-cn"

// MemoryCode memory code
type MemoryCode int

// Tr translate code to description
func (code MemoryCode) Tr(lang string) string {
	if code.IsExist(lang) {
		lang = DefaultLang
	}
	codes := code2Desc[lang]
	desc, ok := codes[code]
	if !ok {
		return fmt.Sprint(code)
	}
	return desc
}

// StatusCode http status code
func (code MemoryCode) StatusCode() int {
	switch code {
	case Success:
		return http.StatusOK

	default:
		return http.StatusBadRequest
	}
}

// IsExist is supported language
func (code MemoryCode) IsExist(lang string) bool {
	for k := range code2Desc {
		if k == lang {
			return true
		}
	}
	return false
}

// response code
var (
	Success MemoryCode = 0

	ErrSystemInternal    MemoryCode = 1000
	ErrBadRequest        MemoryCode = 1001
	ErrInvalidUsername   MemoryCode = 1002
	ErrInvalidPassword   MemoryCode = 1003
	ErrIncorrectPassword MemoryCode = 1004
	ErrNotLoggedIn       MemoryCode = 1005
	ErrNotFoundUser      MemoryCode = 1006
)

var code2Desc = map[string]map[MemoryCode]string{
	"zh-cn": {
		Success: "操作成功",

		ErrSystemInternal:    "系统内部错误，请稍后重试",
		ErrBadRequest:        "请求错误，请检查参数",
		ErrInvalidUsername:   "用户名无效，请输入邮箱",
		ErrInvalidPassword:   "密码无效，至少6位",
		ErrIncorrectPassword: "密码不正确，请核对",
		ErrNotLoggedIn:       "用户未登录，请登录",
		ErrNotFoundUser:      "用户不存在，请注册",
	},
	"en-us": {
		Success: "operation success",

		ErrSystemInternal:    "system internal error，please try again later",
		ErrBadRequest:        "bad request, check params please",
		ErrInvalidUsername:   "invalid username, should be email",
		ErrInvalidPassword:   "invalid passowrd, at least 6 characters",
		ErrIncorrectPassword: "incorrect password，check please",
		ErrNotLoggedIn:       "user not logged in，login please ",
		ErrNotFoundUser:      "user not found, register please",
	},
}
