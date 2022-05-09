package data

type UserLoginOrRegisterData struct {
	UserId int64  `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

type UserInfoData struct {
	User *UserInfo `json:"user" binding:"required"`
}

type UserInfo struct {
	Id            int64  `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	FollowCount   int64  `json:"follow_count" binding:"required"`
	FollowerCount int64  `json:"follower_count" binding:"required"`
	IsFollow      bool   `json:"is_follow" binding:"required"`
}
