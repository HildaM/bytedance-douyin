// Package service
// @Description: follow follower service
// @Author: Quan
package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/repository"
)

type FollowService struct{}
type FollowerService struct{}

//
//  GetFollowList
//  @Description: 获取关注的用户列表
//  @receiver FollowService
//  @param userInfo
//  @return vo.FollowResponseVo
//
func (FollowService) GetFollowList(userInfo vo.FollowListVo) (vo.FollowResponseVo, error) {
	var followedUserList vo.FollowResponseVo

	// TODO 获取该用户的关注列表
	// 	从UserService中获取UserInfoBo的List
	//	根据user_id获取userList
	var userId = userInfo.UserId
	var err error
	followDao := repository.GroupApp.FollowDao
	followedUserList, err = followDao.GetFollowList(userId)
	if err != nil {
		// TODO
	}

	return followedUserList, nil
}
