package service

import (
	"bytedance-douyin/api/vo"
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

func (s VideoService) GetVideoList(userId int64) ([]*vo.Video, error) {
	list, err := videoDao.GetVideoList(userId)
	if err != nil {
		return nil, err
	}
	res := make([]*vo.Video, 0)
	
	author, err := userDao.GetUser(userId)
	
	for _, v := range *list {
		video := &vo.Video{
			Id: v.Id,
			Author: &vo.Author{
				Id:            author.ID,
				Name:          author.Name,
				FollowCount:   author.FollowCount,
				FollowerCount: author.FollowerCount,
				IsFollow:      false,
			},
			// Author:        nil,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
		}
		res = append(res, video)
	}
	return res, nil
}
