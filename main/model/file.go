package model

type File struct {
	Preset    `xorm:"extends"`
	Name      string
	Path      string
	Size      int64
	Hash      string
	OriginID  string
	UserID    int
	FolderID  int
	Thumbnail []byte
	Metadata  map[string]any
}

func (File) TableName() string {
	return "files"
}
