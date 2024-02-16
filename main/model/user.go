package model

import "GoProject/main/enum"

type User struct {
	Preset      `xorm:"extends"`
	Username    string          `xorm:"varchar notnull unique"`
	Email       string          `xorm:"varchar"`
	Passwd      string          `xorm:"varchar notnull"`
	Authority   enum.Authority  `xorm:"smallint notnull default 2"`
	GroupID     uint            `xorm:"notnull default 1"`
	Store       uint64          `xorm:"notnull default 21474836480"`
	TwiceVerify string          `xorm:"varchar"`
	Avatar      string          `xorm:"varchar"`
	Others      *map[string]any `xorm:"text"`
}

func (u User) TableName() string {
	return "users"
}
