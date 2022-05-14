package exceptions

var (
	LoginError     = NewErr("用户名或密码不正确！")
	UserExistError = NewErr("用户名已存在！")
)
