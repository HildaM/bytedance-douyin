package data

type FollowData struct {
	UserList []*UserInfo `json:"user_list" binding:"required"`
}

type FollowerData struct {
	UserList []*UserInfo `json:"user_list" binding:"required"`
}
