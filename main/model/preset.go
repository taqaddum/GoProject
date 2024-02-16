package model

import "time"

type Preset struct {
	ID        int        `xorm:"pk autoincr"`
	CreatedAt time.Time  `xorm:"created timestamp(3)"`
	UpdatedAt time.Time  `xorm:"updated timestamp(3)"`
	DeletedAt *time.Time `xorm:"deleted timestamp(3)"`
}
