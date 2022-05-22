package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
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

func (s VideoService) GetVideoList(userId, myId int64) ([]*vo.Video, error) {
	var videoErr, userErr, followErr error
	var list *[]bo.Video
	var author model.UserDao

	// 获取用户视频列表
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		list, videoErr = videoDao.GetVideoListByUserId(userId)
	}()

	// 获取用户信息
	go func() {
		defer wg.Done()
		author, userErr = userDao.GetUser(userId)
	}()

	// 判断是否是粉丝
	var isFollow bool
	var count int64
	isMyself := userId == myId
	go func() {
		defer wg.Done()
		if isMyself {
			isFollow = false
			return
		}

		count, followErr = GroupApp.FollowService.GetFollowCount(vo.FollowVo{UserId: myId, ToUserId: userId})
		if count != 0 {
			isFollow = true
		}
	}()
	wg.Wait()
	if videoErr != nil {
		return nil, videoErr
	}
	if userErr != nil {
		return nil, userErr
	}
	if followErr != nil {
		return nil, followErr
	}

	res := make([]*vo.Video, 0)

	for _, v := range *list {
		video := &vo.Video{
			Id: v.Id,
			Author: &vo.Author{
				Id:            author.ID,
				Name:          author.Name,
				FollowCount:   author.FollowCount,
				FollowerCount: author.FollowerCount,
				IsFollow:      isFollow,
			},
			// Author:        nil,
			PlayUrl:       utils.VideoUrlPrefix + v.PlayUrl,
			CoverUrl:      utils.ImageUrlPrefix + v.CoverUrl,
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
		wg := sync.WaitGroup{}
		wg.Add(2)

		var err1, err2 error
		// 查询用户是否关注
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
			PlayUrl:       utils.VideoUrlPrefix + video.PlayUrl,
			CoverUrl:      utils.ImageUrlPrefix + video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFollowVideo,
			Title:         video.Title,
			CreatedAt:     video.CreatedAt,
		}

	}

	return videoList, nil
}
