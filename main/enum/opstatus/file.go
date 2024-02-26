package opstatus

type File uint

const (
	UploadFileError File = 10700
	NotFoundFile    File = 10701
)

func (f File) String() string {
	switch f {
	case UploadFileError:
		return "上传文件异常"
	case NotFoundFile:
		return "文件不存在"
	default:
		return "unknown"
	}
}
