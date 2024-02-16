package util

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func BcryptPasswd(pwd string) string {
	pwdByte := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		slog.Debug("加密失败", err.Error())
		return ""
	}
	return string(hash)
}

func CheckPasswd(dbPasswd, reqPasswd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPasswd), []byte(reqPasswd))
	if err != nil {
		slog.Debug("密码校验失败", err.Error())
		return false
	}
	return true
}
