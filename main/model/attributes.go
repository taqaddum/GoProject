package model

type Attributes struct {
	Preset
	Key   string `xorm:"key"`
	Value string `xorm:"value"`
}
