// Package i18n provides ...
package i18n

type translator interface {
	Tr(lang string, code ErrorCode, args ...interface{}) string
}

// language list
const (
	LangZhCN = "zh-cn"
	LangEnUS = "en-us"
)

// i18n i18n instance
type i18n struct {
	defaultLang string
	supported   []string
	translator  translator
}

var i18nInstance i18n

func init() {
	// default supported all language
	i18nInstance = i18n{
		defaultLang: LangZhCN,
		supported:   []string{LangZhCN, LangEnUS},
		translator:  MemoryTranslator{},
	}
}

// GetDefaultLang get default language
func GetDefaultLang() string {
	return i18nInstance.defaultLang
}

// SetDefaultLang set default language
func SetDefaultLang(lang string) {
	i18nInstance.defaultLang = lang
}

// GetSupportedLang get supported language
func GetSupportedLang() []string {
	return i18nInstance.supported
}

// SetSupportedLang set supported language
func SetSupportedLang(langs []string) {
	i18nInstance.supported = langs
}

// SetTranslator set translator
func SetTranslator(trans translator) {
	i18nInstance.translator = trans
}
