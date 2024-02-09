package mapper

import (
	"GoProject/main/model"
	"log/slog"
	"xorm.io/xorm"
)

type UserMapper struct {
	db *xorm.Engine
}

type UserMapApi interface {
	Create(name string) error
	Exist(name string) bool
}

func NewUserMapper(enging *xorm.Engine) *UserMapper {
	return &UserMapper{db: enging}
}

func (mapper UserMapper) Create(name string) error {
	var table = model.Users{
		Uname: name,
	}

	count, err := mapper.db.Insert(table)
	slog.Debug("create user", "name", name, "count", count, "err", err)
	return err
}

func (mapper UserMapper) Exist(name string) bool {
	has, _ := mapper.db.Exist(&model.Users{Uname: name})
	return has
}
