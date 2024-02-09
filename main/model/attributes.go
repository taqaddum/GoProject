package model

type Attributes struct {
	Key   string `xorm:"key"`
	Value string `xorm:"value"`
}
