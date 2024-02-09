package enum

type Permission uint8

const (
	Admin Permission = iota
	Activated
	Guest
	Prohibit
)

func (p Permission) ToString() string {
	switch p {
	case Admin:
		return "管理员"
	case Activated:
		return "用户"
	case Guest:
		return "游客"
	case Prohibit:
		return "黑名单"
	default:
		return "unknown"
	}
}
