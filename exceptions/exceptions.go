package exceptions

var (
	LoginError           = NewErr("用户名或密码不正确！")
	UserExistError       = NewErr("用户名已存在！")
	RejectRequestError   = NewErr("拒绝访问")
	LoginExpired         = NewErr("登录状态已过期")
	ParamValidationError = NewErr("参数校验失败")
)
