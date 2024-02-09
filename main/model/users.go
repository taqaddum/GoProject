package model

import "GoProject/main/enum"

type Users struct {
	Preset
	UserID      int             `xorm:"PK AUTOINCR"`
	Uname       string          `xorm:"VARCHAR"`
	Email       string          `xorm:"VARCHAR"`
	Password    string          `xorm:"VARCHAR"`
	Permission  enum.Permission `xorm:"SMALLINT(1)"`
	GroupID     int             `xorm:"INTEGER"`
	Storage     uint            `xorm:"BIGINT"`
	TwiceVerify string          `xorm:"VARCHAR"`
	Avatar      string          `xorm:"VARCHAR"`
	Options     map[string]any  `xorm:"TEXT"`
}
