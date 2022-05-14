package exceptions

var (
	LoginError    = NewErr("用户名或密码不正确！")
	RegisterError = NewErr("用户名已存在！")
)
