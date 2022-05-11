package vo

import (
	"github.com/golang-jwt/jwt"
	jsoniter "github.com/json-iterator/go"
)

// 注册、登录
type UserVo struct {
	Username string `form:"username" binding:"required,max=32"`
	Password string `form:"password" binding:"required,max=32"`
}

// 用户信息
type UserInfoVo struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserResponseVo struct {
	UserId int64  `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

type UserInfoResponseVo struct {
	User *userInfo `json:"user" binding:"required"`
}

type userInfo struct {
	Id            int64  `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	FollowCount   int64  `json:"follow_count" binding:"required"`
	FollowerCount int64  `json:"follower_count" binding:"required"`
	IsFollow      bool   `json:"is_follow" binding:"required"`
}

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (m *BaseClaims) MarshalBinary() (data []byte, err error) {
	return jsoniter.Marshal(m)
}
