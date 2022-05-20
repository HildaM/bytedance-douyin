package service

import (
	"bytedance-douyin/service/bo"
)

// VideoService
// Author: jtan
// @Description: 增删查 video 的 service
type VideoService struct {
}

func (s VideoService) PostVideo(post bo.VideoPost) error {
	if err := videoDao.PostVideo(post); err != nil {
		return err
	}
	return nil
}
