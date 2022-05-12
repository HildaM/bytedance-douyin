// Package repository
// @Description: Follow Follower Dao
// @Author: Quan

package repository

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
)

type FollowDao struct{}

// GetFollowList get Follow List
func (FollowDao) GetFollowList(userId int64) (vo.FollowResponseVo, error) {
	var followList vo.FollowResponseVo

	//	1. userId --> to_userId list
	toUserIdList, err := getToUserIdList(userId)
	if err != nil {
		return followList, err
	}

	//	2. to_userId list --> user list
	// select user from t_user where id in (?)
	var userDao UserDao
	users, err := userDao.GetUsers(toUserIdList)
	if err != nil {
		return followList, err
	}

	var userList []*vo.UserInfo
	for _, user := range users {
		userInfo := vo.UserInfo{
			Id:            user.ID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}

		userList = append(userList, &userInfo)
	}
	followList.UserList = userList

	return followList, nil
}

//  getToUserIdList use userId to find to_user_id list
func getToUserIdList(userId int64) ([]int64, error) {
	var follows []model.Follow
	var toUserIdList []int64

	// select to_user_id from t_follow where user_id = userId
	if result := global.GVA_DB.Select("to_user_id").Where("user_id = ?", userId).Find(&follows); result.Error != nil {
		return toUserIdList, result.Error
	}

	for _, follow := range follows {
		toUserIdList = append(toUserIdList, follow.ToUserId)
	}
	return toUserIdList, nil
}
