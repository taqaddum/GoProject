package util

import (
	"GoProject/main/component"
	"errors"
)

func ValidateStruct(data ...any) error {
	if len(data) == 0 {
		return errors.New("传入参数不能为空")
	}

	validate := component.NewValidator()

	for iter := range data {
		if err := validate.Struct(iter); err != nil {
			return err
		}
	}
	return nil
}

func ValidateName(name string) error {
	validate := component.NewValidator()
	return validate.Var(name, "username")
}

func ValidatePasswd(passwd string) error {
	return nil
}
