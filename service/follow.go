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

//
//  GetFollowList
//  @Description: 获取关注的用户列表
//  @receiver FollowService
//  @param userInfo
//  @return vo.FollowResponseVo
//
func (FollowService) GetFollowList(userInfo vo.FollowListVo) (vo.FollowResponseVo, error) {
	var followedUserList vo.FollowResponseVo

	//	根据user_id获取userList
	var userId = userInfo.UserId
	var err error
	followedUserList, err = followDao.GetFollowList(userId)
	if err != nil {
		return followedUserList, err
	}

	return followedUserList, nil
}

// FollowOrNot 关注与取消关注
func (FollowService) FollowOrNot(followInfo vo.FollowVo) (string, error) {
	followBo := bo.FollowBo{
		UserId:   followInfo.UserId,
		ToUserId: followInfo.ToUserId,
	}

	// actionType:1 - 关注	actionType:2 - 取消关注
	var err error
	action := followInfo.ActionType
	switch action {
	case "1":
		err = followDao.FollowUser(followBo)
	case "2":
		err = followDao.UnFollowUser(followBo)
	}
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return "", err
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
