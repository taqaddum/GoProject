package opstatus

type Generic interface {
	Operate | Login
	String() string
}

type Operate uint

const (
	Ok            Operate = 200
	Error         Operate = 400
	Unauthorized  Operate = 10301
	InvalidParams Operate = 10302
)

func (c Operate) String() string {
	switch c {
	case Ok:
		return "操作成功"
	case Error:
		return "操作失败"
	case Unauthorized:
		return "访问未授权"
	case InvalidParams:
		return "请求参数错误"
	default:
		return "unknown"
	}
}
