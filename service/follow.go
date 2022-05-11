// Package service
// @Description: follow follower service
// @Author: Quan
package service

import "bytedance-douyin/api/vo"

type FollowService struct{}
type FollowerService struct{}

//
//  GetFollowList
//  @Description: 获取关注的用户列表
//  @receiver FollowService
//  @param userInfo
//  @return vo.FollowResponseVo
//
func (FollowService) GetFollowList(userInfo vo.FollowListVo) vo.FollowResponseVo {
	var followedUserList vo.FollowResponseVo

	// TODO 获取该用户的关注列表

	return followedUserList
}
