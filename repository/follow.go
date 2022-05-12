// Package repository
// @Description: Follow Follower Dao
// @Author: Quan

package repository

import "bytedance-douyin/api/vo"

type FollowDao struct{}

func (FollowDao) GetFollowList(userId int64) (vo.FollowResponseVo, error) {
	var userList vo.FollowResponseVo
	// TODO
	return userList, nil
}
