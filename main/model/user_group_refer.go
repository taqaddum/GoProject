package model

type UserGroupRefer struct {
	User  `xorm:"extends"`
	Group `xorm:"extends"`
}

func (UserGroupRefer) TableName() string {
	return "users"
}
