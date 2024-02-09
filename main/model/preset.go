package model

import "time"

type Preset struct {
	createdAt time.Time  `xorm:"created"`
	updatedAt time.Time  `xorm:"updated"`
	deletedAt *time.Time `xorm:"deleted"`
}
