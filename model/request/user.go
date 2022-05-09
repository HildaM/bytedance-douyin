package request

// 注册、登录
type UserRequest struct {
	Username string `form:"username" binding:"required,max=32"`
	Password string `form:"password" binding:"required,max=32"`
}

// 用户信息
type UserinfoRequest struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}
