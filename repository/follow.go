// Package repository
// @Description: Follow Follower Dao
// @Author: Quan

package repository

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/exceptions"
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
	"context"
	"gorm.io/gorm"
	"strconv"
)

const (
	FAN_SET_KEY    = "follower:user:" // 用户的粉丝
	FOLLOW_SET_KEY = "follow:user:"   // 用户的关注
)

type FollowDao struct{}

// GetFollowList 用户
// Deprecated please use GetFollowListByRedis
func (FollowDao) GetFollowList(userId int64) (vo.FollowResponseVo, error) {
	var followList vo.FollowResponseVo
	var follows []vo.UserInfo

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

// GetFollowListByRedis use redis to refactor 获取用户的关注列表中的用户信息
func (FollowDao) GetFollowListByRedis(userId int64) (vo.FollowResponseVo, error) {
	var followList vo.FollowResponseVo

	userKey := FOLLOW_SET_KEY + strconv.FormatInt(userId, 10)
	rdb := global.GVA_REDIS
	ctx := context.Background()

	// 1. 从redis中获取关注者的id
	res := rdb.SMembers(ctx, userKey)
	if res.Err() != nil {
		return followList, res.Err()
	}
	followerIds := utils.String2Int64(res.Val())

	// 2. 获取关注者
	followers, err := GroupApp.UserDao.GetUsers(followerIds)
	if err != nil {
		return followList, nil
	}

	follows := make([]vo.UserInfo, len(followers))

	for i, u := range followers {
		userInfo := vo.UserInfo{
			Id:            u.ID,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      true,
		}
		follows[i] = userInfo
	}

	followList.UserList = follows
	return followList, nil
}

// GetToUserIdList use userId to find to_user_id list 获取用户关注用户的id
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

// GetToUserIdListByRedis 使用redis重构，获取用户关注用户的id
func (f FollowDao) GetToUserIdListByRedis(userId int64) ([]int64, error) {
	rdb := global.GVA_REDIS
	userKey := FAN_SET_KEY + strconv.FormatInt(userId, 10)
	ctx := context.Background()
	var toUserList []int64

	res := rdb.SMembers(ctx, userKey)
	if res.Err() != nil {
		return []int64{}, res.Err()
	}

	// 没有查询到结果，从数据库中读取并存入redis
	if len(res.Val()) == 0 {
		list, err := f.GetToUserIdList(userId)
		if err != nil {
			return nil, err
		}
		r := rdb.SAdd(ctx, userKey, list)
		if r.Err() != nil {
			return nil, r.Err()
		}
		toUserList = list
	} else {
		toUserList = utils.String2Int64(res.Val())
	}

	return toUserList, nil
}

//	FollowUser insert into t_follow 关注用户
// 1. 如果不存在，直接创建条目
// 2. 如果表中已经存在条目，直接返回即可
// 3. 用户表添加对应用户的follow_count
func (FollowDao) FollowUser(followInfo bo.FollowBo) error {
	// 1. 前置判断
	var follow model.Follow

	tx := global.GVA_DB.Begin()

	result := tx.Debug().Unscoped().Where("user_id = ? AND to_user_id = ?", followInfo.UserId, followInfo.ToUserId).Find(&follow)
	//fmt.Println(result.RowsAffected)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// 已经关注过了
	if result.RowsAffected == 1 && !follow.DeletedAt.Valid {
		return exceptions.RepeatedFollowError
	}

	// 进行关注
	err := tx.Debug().Create(&model.Follow{UserId: followInfo.UserId, ToUserId: followInfo.ToUserId}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新关注和粉丝数
	err = updateFollowAndFollowerCount(tx, followInfo.UserId, followInfo.ToUserId, 1)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func updateFollowAndFollowerCount(tx *gorm.DB, userId, toUserId int64, incr int) error {
	// to_user_id粉丝加一
	err := GroupApp.UserDao.UserFollowerCountIncrement(toUserId, incr)
	if err != nil {
		tx.Rollback()
		return err
	}

	// user_id 关注加一
	err = GroupApp.UserDao.UserFollowCountIncrement(userId, incr)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// FollowUserByRedis use redis to refactor this function
func (FollowDao) FollowUserByRedis(followInfo bo.FollowBo) error {
	rdb := global.GVA_REDIS
	ctx := context.Background()

	userIdInt, toUserIdInt := followInfo.UserId, followInfo.ToUserId
	userId := strconv.FormatInt(userIdInt, 10)
	toUserId := strconv.FormatInt(toUserIdInt, 10)
	//followerKey := FAN_SET_KEY + userId
	//followKey := FOLLOW_SET_KEY + toUserId
	followerKey := FAN_SET_KEY + toUserId
	followKey := FOLLOW_SET_KEY + userId

	// 1. 前置判断，判断用户是否已经关注过to_user_id
	result := rdb.SIsMember(ctx, followKey, toUserId)
	if result.Err() != nil {
		return result.Err()
	}

	// 如果已经关注了，退出即可
	if result.Val() == true {
		return exceptions.RepeatedFollowError
	}

	// 写入数据库
	tx := global.GVA_DB.Begin()
	err := tx.Debug().Create(&model.Follow{UserId: userIdInt, ToUserId: toUserIdInt}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = updateFollowAndFollowerCount(tx, userIdInt, toUserIdInt, 1)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. 关注操作，将关注的信息写入到redis服务器中
	// user关注
	res := rdb.SAdd(ctx, followKey, toUserId)
	if res.Err() != nil || res.Val() <= 0 {
		tx.Rollback()
		return res.Err()
	}

	// 更新对方粉丝的列表
	res = rdb.SAdd(ctx, followerKey, userId)
	if res.Err() != nil || res.Val() <= 0 {
		tx.Rollback()
		return res.Err()
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// UnFollowUser delete row from t_follow
// 如果在已经关注的情况下，存在deleted_at。则先删除deleted_at条目，再将最新的关注标记为”软删除“。已达到更新软删除的目的
func (FollowDao) UnFollowUser(followInfo bo.FollowBo) error {
	// 1. 前置判断
	var follow model.Follow

	tx := global.GVA_DB.Begin()

	result := tx.Debug().Where("user_id = ? and to_user_id = ?", followInfo.UserId, followInfo.ToUserId).Find(&follow)
	//fmt.Println(result.RowsAffected)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// 没关注
	if result.RowsAffected == 0 {
		tx.Rollback()
		return exceptions.UnfollowError
	}
	// 关注了，取消关注
	err := tx.Debug().Delete(&follow).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// to_user_id 粉丝减一
	err = updateFollowAndFollowerCount(tx, followInfo.UserId, followInfo.ToUserId, -1)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// UnFollowUserByRedis use redis to refactor this function
func (FollowDao) UnFollowUserByRedis(followInfo bo.FollowBo) error {
	rdb := global.GVA_REDIS
	ctx := context.Background()

	userId := strconv.FormatInt(followInfo.UserId, 10)
	toUserId := strconv.FormatInt(followInfo.ToUserId, 10)
	followerKey := FAN_SET_KEY + toUserId
	followKey := FOLLOW_SET_KEY + userId

	// 1. 前置判断
	result := rdb.SIsMember(ctx, followerKey, toUserId)
	if result.Err() != nil {
		return result.Err()
	}
	// 没关注
	if result.Val() == false {
		return exceptions.UnfollowError
	}

	var follow model.Follow

	tx := global.GVA_DB.Begin()

	res := tx.Debug().Where("user_id = ? and to_user_id = ?", followInfo.UserId, followInfo.ToUserId).Find(&follow)
	if res.Error != nil {
		return res.Error
	}

	err := tx.Debug().Delete(&follow).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// to_user_id 粉丝减一
	err = updateFollowAndFollowerCount(tx, followInfo.UserId, followInfo.ToUserId, -1)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. 取消关注
	// user取消关注
	r := rdb.SRem(ctx, followKey, toUserId)
	if r.Err() != nil || r.Val() <= 0 {
		return r.Err()
	}

	// 对方粉丝列表移除
	r = rdb.SRem(ctx, followerKey, userId)
	if r.Err() != nil || r.Val() <= 0 {
		return r.Err()
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// GetIsFollow 判断是否关注
func (FollowDao) GetIsFollow(followInfo bo.FollowBo) (bool, error) {
	db := global.GVA_DB

	var count int64
	if err := db.Model(&model.Follow{}).Where("user_id = ? and to_user_id = ?", followInfo.UserId, followInfo.ToUserId).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetIsFollowByRedis 判断是否关注，to_user_id的粉丝中是否有user_id
func (FollowDao) GetIsFollowByRedis(followInfo bo.FollowBo) (bool, error) {
	rdb := global.GVA_REDIS
	followerKey := FAN_SET_KEY + strconv.FormatInt(followInfo.ToUserId, 10)

	res := rdb.SIsMember(context.Background(), followerKey, followInfo.UserId)
	if res.Err() != nil {
		return res.Val(), res.Err()
	}

	return res.Val(), nil
}

// GetFanList 获取粉丝列表
func (FollowDao) GetFanList(userInfo vo.FollowerListVo) (vo.FollowerResponseVo, error) {
	var fansList vo.FollowerResponseVo
	var fans []vo.UserInfo

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

// GetFanListByRedis 获取粉丝列表
// 由于redis中存储的是用户的id，所以获取到id集合后，还需要用id在数据库中查找对应的用户
func (FollowDao) GetFanListByRedis(userInfo vo.FollowerListVo) (vo.FollowerResponseVo, error) {
	rdb := global.GVA_REDIS
	ctx := context.Background()
	var fansList vo.FollowerResponseVo

	// 0. init value
	toUserId := strconv.FormatInt(userInfo.UserId, 10) // 被访问的用户，有可能不是token用户
	userId := strconv.FormatInt(userInfo.TokenId, 10)  // tokenId指代的是当前操作的用户
	followKey := FOLLOW_SET_KEY + userId
	followerKey := FAN_SET_KEY + toUserId

	// 1. 获取指定用户user的粉丝列表
	result := rdb.SMembers(ctx, followerKey)
	if result.Err() != nil {
		return fansList, result.Err()
	}

	fansIds := utils.String2Int64(result.Val())
	userFans, err := GroupApp.UserDao.GetUsers(fansIds)
	if err != nil {
		return fansList, err
	}

	// 2. 判断当前用户是否关注了这些粉丝
	fans := make([]vo.UserInfo, len(userFans))
	for i, u := range userFans {
		// 互关判断
		res := rdb.SIsMember(ctx, followKey, strconv.FormatInt(u.ID, 10))
		if res.Err() != nil {
			return fansList, res.Err()
		}
		isFollow := res.Val()

		fan := vo.UserInfo{
			Id:            u.ID,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      isFollow,
		}

		fans[i] = fan
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

// GetFanCountByRedis 获取粉丝数目
func (FollowDao) GetFanCountByRedis(followInfo bo.FollowBo) (int64, error) {
	rdb := global.GVA_REDIS
	followerKey := FOLLOW_SET_KEY + strconv.FormatInt(followInfo.UserId, 10)

	res := rdb.SCard(context.Background(), followerKey)
	if res.Err() != nil {
		return 0, res.Err()
	}

	return res.Val(), nil
}

// GetFollowUserIdByUserId 根据userId及toUserId判断这些用户是否被关注
func (FollowDao) GetFollowUserIdByUserId(userId int64, toUserIdList []int64) (map[int64]bool, error) {
	isFollowList := make([]int64, 0)

	err := global.GVA_DB.Model(&model.Follow{}).Select("to_user_id").Where("user_id = ?", userId).
		Where("to_user_id IN (?)", toUserIdList).Find(&isFollowList).Error
	if err != nil {
		return nil, err
	}

	followMap := make(map[int64]bool, len(isFollowList))
	for _, v := range isFollowList {
		followMap[v] = true
	}
	return followMap, nil
}

// GetFollowUserIdByUserIdByRedis 根据userId获取关注列表的映射
func (FollowDao) GetFollowUserIdByUserIdByRedis(userId int64, toUserIdList []int64) (map[int64]bool, error) {
	rdb := global.GVA_REDIS
	userKey := FAN_SET_KEY + strconv.FormatInt(userId, 10)

	// SMembersMap：把集合里的元素转换成map的key
	// map[100:{} 200:{} 300:{} 400:{} 500:{} 600:{}]  相当于转换为一个set
	res := rdb.SMembersMap(context.Background(), userKey)
	if res.Err() != nil {
		return map[int64]bool{}, res.Err()
	}

	toUserIdSet := res.Val()
	followMap := make(map[int64]bool, len(toUserIdSet))
	for _, toUserId := range toUserIdList {
		// 如果toUserId存在与user的关注列表中
		// 注意Set的key是string类型！
		if _, ok := toUserIdSet[strconv.FormatInt(toUserId, 10)]; ok == true {
			followMap[toUserId] = true
		}
	}

	return followMap, nil
}
