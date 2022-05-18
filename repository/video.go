// Package repository
// @Description: Video Dao接口
// @Author: Quan

package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
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
