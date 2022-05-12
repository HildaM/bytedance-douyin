package api

import "bytedance-douyin/service"

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:53
 */

type Group struct {
	UserApi
	CommentApi
	LikeApi
	FollowApi
	VideoApi
}

var GroupApp Group

// service
var (
	userService     = service.GroupApp.UserService
	followService   = service.GroupApp.FollowService
	followerService = service.GroupApp.FollowerService
)
