package data

import "bytedance-douyin/api/vo"

type UserLoginOrRegister struct {
	UserId int64  `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

type UserInfo struct {
	User *vo.UserInfo `json:"user" binding:"required"`
}