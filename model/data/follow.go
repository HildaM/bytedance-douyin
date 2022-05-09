package data

type FollowData struct {
	UserList []*UserInfo `json:"user_list"`
}

type FollowerData struct {
	*FollowData
}
