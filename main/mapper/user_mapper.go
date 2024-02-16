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
	Create(user *model.User) error
	Exist(name string) bool
	GetByName(name string) *model.UserGroupRefer
}

func NewUserMapper(engine *xorm.Engine) *UserMapper {
	return &UserMapper{db: engine}
}

func (mapper UserMapper) Create(user *model.User) error {
	_, err := mapper.db.Omit("authority", "store", "group_id").Insert(user)
	return err
}

func (mapper UserMapper) Exist(name string) bool {
	has, _ := mapper.db.Exist(&model.User{Username: name})
	return has
}

func (mapper UserMapper) GetByName(name string) *model.UserGroupRefer {
	var user = &model.UserGroupRefer{}
	has, err := mapper.db.Join("inner", "groups", "users.group_id = groups.id").Where("users.name = ?", name).Get(user)
	if !has {
		return nil
	}

	slog.Debug("get user", "user", user, "err", err)
	return user
}
