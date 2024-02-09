package service

import "GoProject/main/mapper"

type UserService struct {
	mapperApi mapper.UserMapApi
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{mapperApi: userMapper}
}

type UserSrvApi interface {
	LoginByName(username, password string) string
	LoginByPhone(phone string) string
}

func (srv *UserService) LoginByPhone(phone string) string {
	//TODO implement me
	panic("implement me")
}

func (srv *UserService) LoginByName(username string, password string) string {
	if srv.mapperApi.Exist(username) {

	}
	return ""
}
