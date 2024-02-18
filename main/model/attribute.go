package model

type Attributes struct {
	Preset `xorm:"extends"`
	Key    string `xorm:"key"`
	Value  string `xorm:"value"`
}
