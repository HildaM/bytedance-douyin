// Package service
// @Description: follow follower service
// @Author: Quan
package service

import (
	"bytedance-douyin/api/vo"
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
	followedUserList, err = followDao.GetFollowListByRedis(userId)
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
		err = followDao.FollowUserByRedis(followBo)
	case UNFOLLOW:
		err = followDao.UnFollowUserByRedis(followBo)
	}
	if err != nil {
		//global.GVA_LOG.Error(err.Error())
		return 0, err
	}

	return action, nil
}

// GetIsFollow 获取是否关注
func (FollowService) GetIsFollow(followInfo vo.FollowVo) (bool, error) {
	followBo := bo.FollowBo{
		UserId:   followInfo.UserId,
		ToUserId: followInfo.ToUserId,
	}

	//isFollow, err := followDao.GetIsFollow(followBo)
	isFollow, err := followDao.GetIsFollowByRedis(followBo)
	if err != nil {
		//global.GVA_LOG.Error(err.Error())
		return isFollow, err
	}

	return isFollow, nil
}

// GetFanList 获取粉丝列表
func (FollowerService) GetFanList(userInfo vo.FollowerListVo) (vo.FollowerResponseVo, error) {
	var fanList vo.FollowerResponseVo
	var err error

	//fanList, err = followDao.GetFanList(userInfo)
	fanList, err = followDao.GetFanListByRedis(userInfo)
	if err != nil {
		//global.GVA_LOG.Error(err.Error())
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
		//global.GVA_LOG.Error(err.Error())
		return count, err
	}

	return count, nil
}
