package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
)

type LikeDao struct{}

func (LikeDao LikeDao) GetLikeList(userId int64) ([]model.Video, error) {
	var videos []model.Video

	//从 userId获取 video_id list
	videoIds, err := LikeDao.GetVideoIdListByUserId(userId)
	if err != nil {
		return videos, err
	}
	var videoDao VideoDao
	videos, err = videoDao.GetVideos(videoIds)
	if err != nil {
		return videos, err
	}

	return videos, nil

}

func (LikeDao) GetVideoIdListByUserId(userId int64) ([]int64, error) {
	var likes []model.Like
	var videoIdList []int64

	find := global.GVA_DB.Debug().Select("video_id").Where("user_id = ?", userId).Find(&likes)
	if find.Error != nil {
		return videoIdList, find.Error
	}
	//提前指定长度
	videoIdList = make([]int64, 0, len(likes))
	for _, like := range likes {
		videoIdList = append(videoIdList, like.VideoId)
	}
	return videoIdList, nil
}

func (LikeDao) GetVideoLike(videoId int64) (int64, error) {
	db := global.GVA_DB

	var count int64
	if err := db.Model(&model.Like{}).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (LikeDao) GetIsFavorite(videoLikedBo bo.VideoLikedBo) (int64, error) {
	db := global.GVA_DB

	var count int64
	if err := db.Model(&model.Like{}).Where("user_id = ? and video_id = ?", videoLikedBo.UserId, videoLikedBo.VideoId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

//var lock sync.Mutex

// LikeVideo 并发时获赞次数可能会变？加锁？
func (LikeDao) LikeVideo(likedInfo bo.VideoLikedBo) error {

	tx := global.GVA_DB.Begin()
	like := model.Like{
		UserId:  likedInfo.UserId,
		VideoId: likedInfo.VideoId,
	}
	//var videoDao VideoDao
	tx.Debug().Create(&like)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	err := GroupApp.VideoDao.VideoFavoriteCountIncr(likedInfo.VideoId, 1)
	if err != nil {
		tx.Rollback()
		return err
	}

	//lock.Lock()
	////根据video_id查询视频，获得video信息
	//video, err := videoDao.GetVideoById(likedInfo.VideoId)
	//if err != nil {
	//	lock.Unlock()
	//	return err
	//}
	////更新video的favorite_count字段
	//tx.Debug().Model(&video).Update("favorite_count", video.FavoriteCount+1)
	//if tx.Error != nil {
	//	tx.Rollback()
	//	lock.Unlock()
	//	return tx.Error
	//}
	//lock.Unlock()

	//用户若有获赞数量，则根据video author查用户表，再在获赞数量字段+1
	//video中favorite_count字段+1

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (LikeDao) UnLikeVideo(likedInfo bo.VideoLikedBo) error {
	tx := global.GVA_DB.Begin()
	unlike := model.Like{
		UserId:  likedInfo.UserId,
		VideoId: likedInfo.VideoId,
	}
	//var videoDao VideoDao
	//暂时先硬删除
	tx.Debug().Unscoped().Where("user_id = ? and video_id = ?", likedInfo.UserId, likedInfo.VideoId).Delete(&unlike)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	//lock.Lock()
	////根据video_id查询视频，获得video信息
	//video, err := videoDao.GetVideoById(likedInfo.VideoId)
	//if err != nil {
	//	lock.Unlock()
	//	return err
	//}
	////更新video的favorite_count字段
	//tx.Debug().Model(&video).Update("favorite_count", video.FavoriteCount-1)
	//if tx.Error != nil {
	//	tx.Rollback()
	//	lock.Unlock()
	//	return tx.Error
	//}
	//lock.Unlock()

	//用户若有获赞数量，则根据video author查用户表，再在获赞数量字段-1
	//video中favorite_count字段-1
	err := GroupApp.VideoDao.VideoFavoriteCountIncr(likedInfo.VideoId, -1)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return tx.Error
	}

	return nil
}

func (LikeDao) GetFollowVideoIdByUserId(userId int64, videoIdList []int64) (map[int64]bool, error) {
	followList := make([]int64, 0)

	err := global.GVA_DB.Model(&model.Like{}).Select("video_id").Where("user_id = ?", userId).Where("video_id IN (?)", videoIdList).Find(&followList).Error
	if err != nil {
		return nil, err
	}

	followMap := make(map[int64]bool, len(followList))
	for _, v := range followList {
		followMap[v] = true
	}

	return followMap, nil
}
