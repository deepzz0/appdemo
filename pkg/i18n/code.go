// Package i18n provides ...
package i18n

import (
	"net/http"
)

// ErrorCode error code
type ErrorCode int

// StatusCode http status code
func (code ErrorCode) StatusCode() int {
	switch code {
	case Success:
		return http.StatusOK

	default:
		return http.StatusBadRequest
	}
}

// Tr translate code to description
func (code ErrorCode) Tr(lang string, args ...interface{}) string {
	found := false
	for _, l := range i18nInstance.supported {
		if l == lang {
			found = true
			break
		}
	}
	if !found {
		lang = i18nInstance.defaultLang
	}
	return i18nInstance.translator.Tr(lang, code, args...)
}

// response code
var (
	Success ErrorCode = 0

	ErrSystemInternal    ErrorCode = 1000
	ErrBadRequest        ErrorCode = 1001
	ErrInvalidUsername   ErrorCode = 1002
	ErrInvalidPassword   ErrorCode = 1003
	ErrIncorrectPassword ErrorCode = 1004
	ErrNotLoggedIn       ErrorCode = 1005
	ErrNotFoundUser      ErrorCode = 1006
	ErrAlreadyExistUser  ErrorCode = 1007
)
