package myValidator

import (
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

		//自定义规则翻译
		err := v.RegisterTranslation("alphanum", trans, func(ut ut.Translator) error {
			return ut.Add("alphanum", "{0}只能包含字母和数字", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("alphanum", fe.Field())
			return t
		})
		if err != nil {
			return
		}
	}
}

func GetTrans() ut.Translator {
	return trans
}
