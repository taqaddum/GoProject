package opstatus

type Login uint

const (
	UserNotFound Login = iota + 2001
	WrongPassword
	UnLogin
	TokenExpired
)

func (l Login) String() string {
	switch l {
	case UserNotFound:
		return "用户不存在"
	case WrongPassword:
		return "用户名或密码错误"
	case UnLogin:
		return "未登录"
	case TokenExpired:
		return "token已过期"
	default:
		return "Unknown error"
	}
}
