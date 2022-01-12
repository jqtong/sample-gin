package request

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zht "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)


func init() {

}

// Validate common validator
func Validate(obj interface{}) ([]string, error) {

	enl := en.New()
	zhl := zh.New()

	v := validator.New()
	uni := ut.New(zhl, zhl, enl)

	translator, _ := uni.GetTranslator("zh")

	if err := zht.RegisterDefaultTranslations(v, translator); err != nil {
		return nil, err
	}

	//获取 comment 作为翻译字段
	zht.RegisterDefaultTranslations(v, translator)
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

	//检查手机号是否符合规则
	_ = v.RegisterValidation("checkPhoneRex", checkPhoneRex)
	_ = v.RegisterTranslation("checkPhoneRex", translator, func(ut ut.Translator) error {
		return ut.Add("checkPhoneRex", "{0} 不符合规则", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("checkPhoneRex", fe.Field(), fe.Value().(string))
		return t
	})

	errors := make([]string, 0)

	err := v.Struct(obj)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errs {
				errors = append(errors, err.Translate(translator))
			}
		}

		return errors, err
	}

	return nil, nil
}

func checkPhoneRex(fl validator.FieldLevel) bool {
	// 获取手机号
	phone := fl.Field().String()
	regular := "^1(3\\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}
