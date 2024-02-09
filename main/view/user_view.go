package view

import "GoProject/main/enum"

type UserBaseInfo struct {
	Id       int             `json:"id"`
	Username string          `json:"username" form:"username"`
	Password string          `json:"-" form:"password"`
	Email    string          `json:"email" form:"email"`
	Group    string          `json:"group"`
	Storage  uint            `json:"storage"`
	Avatar   []byte          `json:"avatar"`
	Role     enum.Permission `json:"role"`
}
