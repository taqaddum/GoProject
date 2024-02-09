package model

type Downloads struct {
	Preset
	TaskID     int            `xorm:"PK AUTOINCR"`
	Status     uint8          `xorm:"SMALLINT(1)"`
	OriginID   string         `xorm:"TEXT"`
	Kind       uint8          `xorm:"SMALLINT(1)"`
	TotalSize  uint           `xorm:"BIGINT"`
	ActualSize uint           `xorm:"BIGINT"`
	Speed      uint           `xorm:"BIGINT"`
	UserID     int            `xorm:"INT"`
	Attributes map[string]any `xorm:"TEXT"`
	ErrInfo    string         `xorm:"TEXT"`
	HostDir    string         `xorm:"VARCHAR"`
	FolderID   int            `xorm:"INT"`
}
