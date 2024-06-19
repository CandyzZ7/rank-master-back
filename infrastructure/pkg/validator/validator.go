package validator

import (
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func New() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		_ = validate.RegisterValidation("phone", validatePhone)

	})
	return validate
}

func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if phone == "" {
		return true
	}
	matched, _ := regexp.MatchString(`^(?:(?:\+|00)86)?1[3-9]\d{9}$`, phone)
	return matched // 验证通过
}
