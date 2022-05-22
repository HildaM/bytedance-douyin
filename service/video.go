package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/service/bo"
	"sync"
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
	list, err := videoDao.GetVideoListByUserId(userId)
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

// GetVideoFeed 获取视频流，并查询follow及favorite更新is_follow字段
func (s VideoService) GetVideoFeed(userId, t int64) ([]vo.Video, error) {

	videos, err := videoDao.GetVideoListByTime(t)
	if err != nil {
		return nil, err
	}

	toUserIdList := make([]int64, len(videos))
	toVideoIdList := make([]int64, len(videos))
	for i, v := range videos {
		toUserIdList[i] = v.AuthorId
		toVideoIdList[i] = v.ID
	}

	// 查询视频的发布者是否是用户的关注
	// 如果没有登录，则全部没关注
	// 否则，通过查询数据库确定
	followUserMap := make(map[int64]bool, len(toUserIdList))
	followVideoMap := make(map[int64]bool, len(toVideoIdList))
	// 没有登录
	if userId == 0 {
		// 全部赋值为false
		for _, uId := range toUserIdList {
			followUserMap[uId] = false
			followVideoMap[uId] = false
		}
	} else {
		// 查询用户是否关注
		wg := sync.WaitGroup{}
		wg.Add(2)

		var err1, err2 error
		go func() {
			defer wg.Done()
			followUserMap, err1 = followDao.GetFollowUserIdByUserId(userId, toUserIdList)
		}()

		// 查询视频是否点赞
		go func() {
			defer wg.Done()
			followVideoMap, err2 = likeDao.GetFollowVideoIdByUserId(userId, toVideoIdList)
		}()
		wg.Wait()
		if err1 != nil {
			return nil, err1
		}
		if err2 != nil {
			return nil, err2
		}
	}

	videoList := make([]vo.Video, len(videos))

	for i, video := range videos {
		isFollowAuthor, ok := followUserMap[video.AuthorId]
		isFollowVideo, ok2 := followVideoMap[video.ID]
		if !ok {
			isFollowAuthor = false
		}
		if !ok2 {
			isFollowVideo = false
		}
		videoList[i] = vo.Video{Id: video.ID,
			Author: &vo.Author{
				Id:            video.AuthorId,
				Name:          video.Author.Name,
				FollowCount:   video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,
				IsFollow:      isFollowAuthor,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFollowVideo,
			Title:         video.Title,
			CreatedAt:     video.CreatedAt,
		}

	}

	return videoList, nil
}
