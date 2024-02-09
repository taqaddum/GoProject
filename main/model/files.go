package model

type Files struct {
	Preset
	FileID    int
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
