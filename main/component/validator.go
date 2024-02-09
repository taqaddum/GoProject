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

	})
	if err := validate.RegisterValidation("username", nameTag); err != nil {
		slog.Warn("nameTage注册失败", "error", err)
	}
	return validate
}

func nameTag(value validator.FieldLevel) bool {
	name := value.Field().String()
	re, err := regexp.Compile(`^[a-zA-Z0-9_ ]+$`)
	if err != nil {
		slog.Warn("正则表达式编译失败")
		return false
	}
	return re.MatchString(name)
}
