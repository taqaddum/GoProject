package enum

type Authority uint8

const (
	Admin Authority = iota + 1
	Normal
	Guest
	Prohibit
)

func (p Authority) ToString() string {
	switch p {
	case Admin:
		return "管理员"
	case Normal:
		return "普通用户"
	case Guest:
		return "游客"
	case Prohibit:
		return "黑名单"
	default:
		return "unknown"
	}
}
