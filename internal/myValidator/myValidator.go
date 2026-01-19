package myValidator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhCn := zh.New()
		uni := ut.New(zhCn, zhCn)

		trans, _ = uni.GetTranslator("zh")
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)

		//字段翻译
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("label")
			if label != "" {
				return label
			}

			// fallback：没 label 再用 json
			name := field.Tag.Get("json")
			if name == "-" {
				return ""
			}
			return strings.Split(name, ",")[0]
		})

	}
}

func GetTrans() ut.Translator {
	return trans
}
