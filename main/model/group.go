package model

type Group struct {
	Preset `xorm:"extends"`
	Name   string `xorm:"varchar notnull"`
}

func (g *Group) TableName() string {
	return "groups"
}
