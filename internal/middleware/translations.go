package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(context *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := context.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		validate, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
				break
			case "en":
				_ = enTranslations.RegisterDefaultTranslations(validate, trans)
				break
			default:
				_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
				break
			}
			context.Set("trans", trans)
		}

		context.Next()
	}
}
