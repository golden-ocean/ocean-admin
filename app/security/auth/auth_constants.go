package auth

const (
	LoginSuccess  = "登录成功！"
	LoginFail     = "登录失败！"
	LogoutSuccess = "注销成功！"
	LogoutFail    = "注销失败！"

	ErrorUsernameOrPassword = "用户名或密码错误！"
	ErrorDisableStatus      = "账户锁定, 请联系管理员"
	ErrorAccessExpiredTime  = "签名过期, 请重新登陆"
	ErrorCheckExpiredTime   = "签名错误, 检查token是否正确"
	ErrorRefreshExpiredTime = "刷新签名过期, 请重新登陆"
)
