// Package repository
// @Description: Video Dao接口
// @Author: Quan

package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"gorm.io/gorm"
	"time"
)

type VideoDao struct{}

func (VideoDao) GetVideos(videoIds []int64) ([]model.Video, error) {
	var videos []model.Video
	err := global.GVA_DB.Model(&model.Video{}).Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Table(model.UserDao{}.TableName())
	}).Where("id in (?)", videoIds).Find(&videos).Error
	if err != nil {
		return videos, err
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
		Title:         post.Title,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	if result := global.GVA_DB.Create(&video); result.Error != nil {
		return result.Error
	}
	return nil
}

func (VideoDao) GetVideoListByUserId(userId int64) ([]model.Video, error) {
	videos := make([]model.Video, 0)
	result := global.GVA_DB.Model(&model.Video{}).Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Table(model.UserDao{}.TableName())
	}).Where("author_id = ?", userId).Find(&videos)

	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

func (VideoDao) GetVideoListByTime(t int64) ([]model.Video, error) {
	videos := make([]model.Video, 0)
	//var video model.Video

	// 查询出最新的30个视频
	err := global.GVA_DB.Model(&model.Video{}).Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Table(model.UserDao{}.TableName())
	}).Where("created_at <= ?", time.UnixMilli(t).Format("2006-01-02 15:04:05")).
		Order("created_at desc").
		Limit(30).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (VideoDao) VideoFavoriteCountIncr(videoId int64, incr int) error {
	var video model.Video

	err := global.GVA_DB.Model(&video).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", incr)).Error
	if err != nil {
		return err
	}
	return nil
}

func (VideoDao) VideoCommentCountIncr(videoId int64, incr int) error {
	var video model.Video

	err := global.GVA_DB.Model(&video).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", incr)).Error
	if err != nil {
		return err
	}
	return nil
}
