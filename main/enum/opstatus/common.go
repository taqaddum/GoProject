package opstatus

type Generic interface {
	Common | Login
	String() string
}

type Common uint

const (
	Ok            Common = 200
	Error         Common = 400
	Unauthorized  Common = 10301
	InvalidParams Common = 10302
)

func (c Common) String() string {
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
