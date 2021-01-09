// Package i18n provides ...
package i18n

import (
	"fmt"
)

// MemoryTranslator memory translator
type MemoryTranslator struct{}

// Tr translate lang
func (trans MemoryTranslator) Tr(lang string, code ErrorCode,
	args ...interface{}) string {

	codes := code2Desc[lang]
	desc, ok := codes[code]
	if !ok {
		return fmt.Sprint(code)
	}
	return fmt.Sprintf(desc, args...)
}

var code2Desc = map[string]map[ErrorCode]string{
	"zh-cn": {
		Success: "操作成功",

		ErrSystemInternal:    "系统内部错误，请稍后重试",
		ErrBadRequest:        "请求错误，请检查参数",
		ErrInvalidUsername:   "用户名无效，请输入邮箱",
		ErrInvalidPassword:   "密码无效，至少6位",
		ErrIncorrectPassword: "密码不正确，请核对",
		ErrNotLoggedIn:       "用户未登录，请登录",
		ErrNotFoundUser:      "用户不存在，请注册",
		ErrAlreadyExistUser:  "用户已存在，请登录",
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
		ErrAlreadyExistUser:  "user already exist, login please",
	},
}
