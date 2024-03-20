package dto

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"

	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var (
	Validate *validator.Validate
	Uni      *ut.UniversalTranslator
)

// 初始化过滤器
func InitValidate() {
	Uni = ut.New(zh2.New())
	Validate = validator.New()
	_ = Validate.RegisterValidation("required_if", RequiredIf)
}

func BindingValidParams(c *gin.Context, param interface{}, _ reflect.Type) (err error) {
	trans, _ := Uni.GetTranslator("zh")
	if err = Validate.Struct(param); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return
}

func RequiredIf(fl validator.FieldLevel) bool {
	paramAValue := fl.Parent().FieldByName("Bot").String()

	for _, v := range strings.Split(fl.Param(), " ") {
		if paramAValue == v {
			val := fl.Field().Interface()

			switch val.(type) {
			case string:
				return val != ""
			case float64:
				return val != 0.0
			case int:
				return val != 0
			default:
				return val != nil
			}
		}
	}
	return true
}
