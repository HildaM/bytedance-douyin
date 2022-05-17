// Package repository
// @Description: Follow Follower Dao
// @Author: Quan

package repository

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"errors"
	"reflect"
)

type FollowDao struct{}

//Deprecated
func (dao FollowDao) GetFollowList2(userId int64) (vo.FollowResponseVo, error) {
	var followList vo.FollowResponseVo

	//	1. userId --> to_userId list
	toUserIdList, err := dao.GetToUserIdList(userId)
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
			IsFollow:      true,
		}

		userList = append(userList, &userInfo)
	}
	followList.UserList = userList

	return followList, nil
}

// GetFollowList pass
func (FollowDao) GetFollowList(userId int64) (vo.FollowResponseVo, error) {
	var followList vo.FollowResponseVo
	var follows []*vo.UserInfo

	err := global.GVA_DB.Raw(
		"SELECT a.to_user_id as id, u.name, u.follow_count, u.follower_count, true as `is_follow`"+
			"FROM (SELECT to_user_id FROM t_follow f WHERE f.user_id = ? and f.deleted_at IS NULL) a"+
			"		LEFT JOIN t_user u ON u.id = a.to_user_id",
		userId,
	).Scan(&follows).Error
	if err != nil {
		return followList, err
	}

	followList.UserList = follows
	return followList, nil
}

// GetToUserIdList use userId to find to_user_id list
func (FollowDao) GetToUserIdList(userId int64) ([]int64, error) {
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

//	FollowUser insert into t_follow
// 1. 如果不存在，直接创建条目
// 2. 如果表中已经存在条目，直接返回即可
func (FollowDao) FollowUser(followInfo bo.FollowBo) error {
	// 1. 前置判断
	var follow model.Follow
	global.GVA_DB.Where("user_id = ? and to_user_id = ?", followInfo.UserId, followInfo.ToUserId).Find(&follow)
	if !reflect.DeepEqual(follow, model.Follow{}) {
		return errors.New("已经关注过了，请勿重复操作")
	}

	// 2. 创建条目
	if err := followUser(followInfo); err != nil {
		return err
	}
	return nil
}

// followUser 第一次关注操作
func followUser(followInfo bo.FollowBo) error {
	tx := global.GVA_DB.Begin()
	follow := model.Follow{
		UserId:   followInfo.UserId,
		ToUserId: followInfo.ToUserId,
	}

	tx.Debug().Create(&follow)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return tx.Error
	}

	return nil
}

// UnFollowUser delete row from t_follow
// 如果在已经关注的情况下，存在deleted_at。则先删除deleted_at条目，再将最新的关注标记为”软删除“。已达到更新软删除的目的
// TODO 漏洞：如果在未关注情况下，连续调用两次这个方法，那么会将最后一个软删除删掉
func (FollowDao) UnFollowUser(followInfo bo.FollowBo) error {
	// 1. 前置判断
	var follow model.Follow
	global.GVA_DB.Unscoped().Where("user_id = ? and to_user_id = ? and deleted_at IS NOT NULL", followInfo.UserId, followInfo.ToUserId).Delete(&follow)

	// 2. 取消关注
	if err := unFollowUser(followInfo); err != nil {
		return err
	}

	return nil
}

func unFollowUser(followInfo bo.FollowBo) error {
	tx := global.GVA_DB.Begin()

	unFollow := model.Follow{
		UserId:   followInfo.UserId,
		ToUserId: followInfo.ToUserId,
	}

	tx.Debug().Where("user_id = ? and to_user_id = ?", followInfo.UserId, followInfo.ToUserId).Delete(&unFollow)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return tx.Error
	}

	return nil
}

// GetFollowCount judge user is following another one
func (FollowDao) GetFollowCount(followInfo bo.FollowBo) (int64, error) {
	db := global.GVA_DB

	var count int64
	if err := db.Model(&model.Follow{}).Where("user_id = ? and to_user_id = ?", followInfo.UserId, followInfo.ToUserId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetFanList 获取粉丝列表
func (FollowDao) GetFanList(userInfo vo.FollowerListVo) (vo.FollowerResponseVo, error) {
	var fansList vo.FollowerResponseVo
	var fans []*vo.UserInfo

	userId := userInfo.UserId
	tokenId := userInfo.TokenId

	err := global.GVA_DB.Raw(
		"SELECT a.user_id as id, u.name, u.follow_count, u.follower_count,"+
			"CASE WHEN a.user_id = b.to_user_id THEN true ELSE false END as `is_follow`"+
			"FROM (SELECT user_id FROM t_follow f WHERE f.to_user_id = ? AND f.deleted_at is NULL) a"+
			"	LEFT JOIN t_follow b ON b.user_id = ? AND a.user_id = b.to_user_id AND b.deleted_at is NULL"+
			"	LEFT JOIN t_user u ON u.id = a.user_id",
		userId, tokenId,
	).Scan(&fans).Error
	if err != nil {
		return fansList, err
	}

	fansList.UserList = fans
	return fansList, nil
}

// GetFanCount 获取粉丝数目
func (FollowDao) GetFanCount(followInfo bo.FollowBo) (int64, error) {
	db := global.GVA_DB

	var count int64
	if err := db.Model(&model.Follow{}).Where("to_user_id = ?", followInfo.UserId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
