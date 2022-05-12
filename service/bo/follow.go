// Package bo
// @Description: follow follower bo
// @Author: Quan

package bo

type FollowListBo struct {
	UserId int64
	Token  string
}

type FollowBo struct {
	UserId   int64
	ToUserId int64
}
