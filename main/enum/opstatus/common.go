package opstatus

type Common uint

const (
	Ok    Common = 200
	Error Common = 400
)

func (c Common) String() string {
	switch c {
	case Ok:
		return "操作成功"
	case Error:
		return "操作失败"
	default:
		return "unknown"
	}
}
