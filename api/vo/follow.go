package vo

// 关注
type FollowVo struct {
	UserId     int64  `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   string `form:"to_user_id" binding:"required"`
	ActionType string `form:"action_type" binding:"required,oneof=1 2"`
}

// 关注列表
type FollowListVo struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// 粉丝列表
type FollowerListVo struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type FollowResponseVo struct {
	UserList []*userInfo `json:"user_list" binding:"required"`
}

type FollowerResponseVo struct {
	UserList []*userInfo `json:"user_list" binding:"required"`
}
