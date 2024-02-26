package component

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"regexp"
	"sync"
)

var onceValidator sync.Once
var validate *validator.Validate

func NewValidator() *validator.Validate {
	onceValidator.Do(func() {
		validate = validator.New()
		if err := validate.RegisterValidation("username", nameTag); err != nil {
			slog.Warn("nameTage注册失败", "error", err)
		}

		if err := validate.RegisterValidation("password", passwordTag); err != nil {
			slog.Warn("passwordTag注册失败", "error", err)
		}
	})

	return validate
}

func nameTag(f validator.FieldLevel) bool {
	name := f.Field().String()
	re, err := regexp.Compile(`^[a-zA-Z0-9_ ]+$`)
	notAllBlank, _ := regexp.MatchString(`^\S+$`, name)
	if err != nil {
		slog.Warn("正则表达式编译失败")
		return false
	}
	return re.MatchString(name) && notAllBlank
}

func passwordTag(f validator.FieldLevel) bool {
	password := f.Field().String()
	if len(password) < 8 {
		return false
	}

	hasLetters, _ := regexp.MatchString(`[A-Za-z]`, password)
	hasDigit, _ := regexp.MatchString(`[0-9]`, password)
	hasSpecialChar, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, password)
	hasSpace, _ := regexp.MatchString(`\s`, password)
	return hasLetters && hasDigit && !hasSpace || hasSpecialChar
}
