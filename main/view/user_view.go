package view

type RegisterReq struct {
	Username        string `json:"username" form:"username" validate:"username,required"`
	Password        string `json:"password" form:"password" validate:"password,required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"eqfield=Password,required"`
	Email           string `json:"email" form:"email" validate:"email,required"`
}

type UserDetailResp struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	Store     uint64 `json:"store"`
	Avatar    string `json:"avatar"`
	GroupName string `json:"group_name"`
}
