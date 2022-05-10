package request

// 关注
type FollowRequest struct {
	UserId     int64  `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   string `form:"to_user_id" binding:"required"`
	ActionType string `form:"action_type" binding:"required,oneof=1 2"`
}

// 关注列表
type FollowListRequest struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// 粉丝列表
type FollowerListRequest struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}
