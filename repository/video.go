// Package repository
// @Description: Video Dao接口
// @Author: Quan

package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"gorm.io/gorm"
)

type VideoDao struct{}

func (VideoDao) GetVideos(videoIds []int64) ([]model.Video, error) {
	var videos []model.Video
	if result := global.GVA_DB.Where("id in (?)", videoIds).Find(&videos); result.Error != nil {
		return videos, result.Error
	}
	return videos, nil
}

func (VideoDao) GetVideoById(videoId int64) (model.Video, error) {
	var video model.Video
	if result := global.GVA_DB.Where("id = ?", videoId).First(&video); result.Error != nil {
		return video, result.Error
	}
	return video, nil
}

// PostVideo
// @Description: 新增视频记录的 Dao
// @Author: jtan
func (VideoDao) PostVideo(post bo.VideoPost) error {
	video := model.Video{
		AuthorId:      post.AuthorId,
		PlayUrl:       post.PlayUrl,
		CoverUrl:      post.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	if result := global.GVA_DB.Create(&video); result.Error != nil {
		return result.Error
	}
	return nil
}

func (VideoDao) GetVideoListByUserId(userId int64) (*[]bo.Video, error) {
	videos := make([]bo.Video, 0)
	result := global.GVA_DB.Model(&model.Video{}).Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Table(model.UserDao{}.TableName())
	}).Where("author_id = ?", userId).Find(&videos)

	if result.Error != nil {
		return nil, result.Error
	}
	return &videos, nil
}

func (VideoDao) GetVideoListByTime(t int64) ([]model.Video, error) {
	videos := make([]model.Video, 0)
	//var video model.Video

	// 查询出最新的30个视频
	err := global.GVA_DB.Model(&model.Video{}).Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Table(model.UserDao{}.TableName())
	}).Where("created_at <= FROM_UNIXTIME(?)", t).Order("created_at desc").
		Limit(30).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}
