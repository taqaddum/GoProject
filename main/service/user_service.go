package service

import (
	"GoProject/main/mapper"
	"GoProject/main/model"
	"GoProject/main/util"
	"GoProject/main/view"
	"time"
)

type UserService struct {
	mapApi mapper.UserMapApi
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{mapApi: userMapper}
}

type UserSrvApi interface {
	Login(username, password string) string
	Register(user *model.User) error
	GetDetail(name string) *view.UserDetailResp
}

func (srv UserService) Login(username string, password string) string {
	user := srv.mapApi.GetByName(username)
	if user != nil && util.CheckPasswd(user.Passwd, password) {
		return util.SignToken(user.ID, user.Username, user.Authority)
	}
	return ""
}

func (srv UserService) Register(user *model.User) error {
	return srv.mapApi.Create(user)
}

func (srv UserService) GetDetail(name string) *view.UserDetailResp {
	user := srv.mapApi.GetByName(name)
	if user == nil {
		return nil
	}

	diff := time.Now().Sub(user.CreatedAt)
	return &view.UserDetailResp{
		Username:  user.Username,
		Email:     user.Email,
		Age:       util.FormatTimeDiff(diff),
		Store:     user.Store,
		Avatar:    user.Avatar,
		GroupName: user.Group.Name,
	}
}
