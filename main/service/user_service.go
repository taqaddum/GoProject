package service

import (
	"GoProject/main/mapper"
	"GoProject/main/model"
	"GoProject/main/util"
	"GoProject/main/view"
	"time"
)

type UserService struct {
	mapperApi mapper.UserMapApi
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{mapperApi: userMapper}
}

type UserSrvApi interface {
	LoginByName(username, password string) string
	LoginByPhone(phone string) string
	Register(user *model.User) error
	GetDetail(name string) *view.UserDetailResp
}

func (srv *UserService) LoginByPhone(phone string) string {
	//TODO implement me
	panic("implement me")
}

func (srv *UserService) LoginByName(username string, password string) string {
	user := srv.mapperApi.GetByName(username)
	if user != nil && util.CheckPasswd(user.Passwd, password) {
		return util.SignToken(user.ID, user.Username, user.Authority)
	}
	return ""
}

func (srv *UserService) Register(user *model.User) error {
	return srv.mapperApi.Create(user)
}

func (srv *UserService) GetDetail(name string) *view.UserDetailResp {
	user := srv.mapperApi.GetByName(name)
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
