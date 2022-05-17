package vo

// FollowVo 关注
type FollowVo struct {
	UserId     int64  `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   int64  `form:"to_user_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required,oneof=1 2"`
}

// FollowListVo 关注列表
type FollowListVo struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// FollowerListVo 粉丝列表
type FollowerListVo struct {
	UserId  int64  `form:"user_id" binding:"required"`
	Token   string `form:"token" binding:"required"`
	TokenId int64
}

type FollowResponseVo struct {
	UserList []*UserInfo `json:"user_list" binding:"required"`
}

type FollowerResponseVo struct {
	UserList []*UserInfo `json:"user_list" binding:"required"`
}
