package configs

import (
	"reflect"

	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Translator ut.Translator
	Validate   *validator.Validate
)

func ValidatorInit(locale string) {
	Validate = validator.New()

	zhCn := zh.New()
	enUs := en.New()
	uni := ut.New(enUs, zhCn, enUs)

	Translator, _ = uni.GetTranslator(locale)

	switch locale {
	case "en":
		_ = en_translations.RegisterDefaultTranslations(Validate, Translator)
	case "zh":
		Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("zh")
			if label == "" {
				return field.Name
			}
			return label
		})
		_ = zh_translations.RegisterDefaultTranslations(Validate, Translator)
	default:
		_ = en_translations.RegisterDefaultTranslations(Validate, Translator)
	}
}

func ValidatorErrors(err error) []string {
	var fields []string
	for _, err := range err.(validator.ValidationErrors) {
		fields = append(fields, err.Translate(Translator))
	}
	return fields
}
