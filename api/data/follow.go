package data

import "bytedance-douyin/api/vo"

type Follow struct {
	UserList []*vo.UserInfo `json:"user_list" binding:"required"`
}

type Follower struct {
	UserList []*vo.UserInfo `json:"user_list" binding:"required"`
}
