package model

type File struct {
	Preset    `xorm:"extends"`
	Name      string
	Path      string
	Size      int64
	Md5       string
	OriginID  string
	UserID    int
	FolderID  int
	Thumbnail []byte
	Metadata  map[string]any
}

func (File) TableName() string {
	return "files"
}
