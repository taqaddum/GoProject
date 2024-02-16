package util

import (
	"GoProject/main/component"
	"errors"
)

var validate = component.NewValidator()

func ValidateStruct(data ...any) error {
	if len(data) == 0 {
		return errors.New("传入参数不能为空")
	}

	for iter := range data {
		if err := validate.Struct(iter); err != nil {
			return err
		}
	}
	return nil
}

func ValidateName(name string) error {
	return validate.Var(name, "username")
}

func ValidatePasswd(passwd string) error {
	return validate.Var(passwd, "password")
}
