package opstatus

type File uint

const (
	UploadFileError File = 10700
)

func (f File) String() string {
	switch f {
	case UploadFileError:
		return "上传文件异常"
	default:
		return "unknown"
	}
}
