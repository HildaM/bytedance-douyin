// Package service
// @Description: follow follower service
// @Author: Quan
package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"bytedance-douyin/service/bo"
)

type FollowService struct{}
type FollowerService struct{}

const (
	FOLLOW   = 1
	UNFOLLOW = 2
)

// GetFollowList
//  @Description: 获取关注的用户列表
//  @receiver FollowService
//  @param userInfo
//  @return vo.FollowResponseVo
//
func (FollowService) GetFollowList(userInfo vo.FollowListVo) (vo.FollowResponseVo, error) {
	var followedUserList vo.FollowResponseVo

	//	根据user_id获取userList
	userId := userInfo.UserId
	var err error
	followedUserList, err = followDao.GetFollowList(userId)
	if err != nil {
		return followedUserList, err
	}

	return followedUserList, nil
}

// FollowOrNot 关注与取消关注
func (FollowService) FollowOrNot(followInfo vo.FollowVo) (int8, error) {
	followBo := bo.FollowBo{
		UserId:   followInfo.UserId,
		ToUserId: followInfo.ToUserId,
	}

	// actionType:1 - 关注	actionType:2 - 取消关注
	var err error
	action := followInfo.ActionType
	switch action {
	case FOLLOW:
		err = followDao.FollowUser(followBo)
	case UNFOLLOW:
		err = followDao.UnFollowUser(followBo)
	}
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return 0, err
	}

	return action, nil
}

// GetFollowCount 获取关注数
func (FollowService) GetFollowCount(followInfo vo.FollowVo) (int64, error) {
	followBo := bo.FollowBo{
		UserId:   followInfo.UserId,
		ToUserId: followInfo.ToUserId,
	}

	count, err := followDao.GetFollowCount(followBo)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return 0, err
	}

	return count, nil
}

// GetFanList 获取粉丝列表
func (FollowerService) GetFanList(userInfo vo.FollowerListVo) (vo.FollowerResponseVo, error) {
	var fanList vo.FollowerResponseVo
	userId := userInfo.UserId
	var err error

	fanList, err = followDao.GetFanList(userId)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return fanList, err
	}

	return fanList, nil
}

// GetFanCount 获取粉丝数
func (FollowerService) GetFanCount(followInfo vo.FollowVo) (int64, error) {
	followBo := bo.FollowBo{
		UserId:   followInfo.UserId,
		ToUserId: followInfo.ToUserId,
	}

	count, err := followDao.GetFanCount(followBo)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return count, err
	}

	return count, nil
}
