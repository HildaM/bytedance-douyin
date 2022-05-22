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

// GetVideoList 获取视频列表，根据user_id及token判断是否是自己的视频列表
// 如果不是则判断是否是视频的点赞者或视频发布用户的粉丝
func (VideoService) GetVideoList(userId, myId int64) ([]vo.Video, error) {

	// 获取用户视频列表
	videoList, videoErr := videoDao.GetVideoListByUserId(userId)
	if videoErr != nil {
		return nil, videoErr
	}

	res, err := handleFollowCondition(videoList, myId)
	if err != nil {
		return nil, err
	}
	return res, nil
	// 判断是否是粉丝
	//var isFollow bool
	//var count int64
	//isMyself := userId == myId

	//go func() {
	//	defer wg.Done()
	//	if isMyself {
	//		isFollow = false
	//		return
	//	}
	//
	//	count, followErr = GroupApp.FollowService.GetFollowCount(vo.FollowVo{UserId: myId, ToUserId: userId})
	//	if count != 0 {
	//		isFollow = true
	//	}
	//}()
	//wg.Wait()
	//if videoErr != nil {
	//	return nil, videoErr
	//}
	//if userErr != nil {
	//	return nil, userErr
	//}
	//if followErr != nil {
	//	return nil, followErr
	//}
	//
	//res := make([]vo.Video, 0)
	//
	//for _, v := range videoList {
	//	video := vo.Video{
	//		Id: v.ID,
	//		Author: &vo.Author{
	//			Id:            v.Author.ID,
	//			Name:          v.Author.Name,
	//			FollowCount:   v.Author.FollowCount,
	//			FollowerCount: v.Author.FollowerCount,
	//			IsFollow:      isFollow,
	//		},
	//		// Author:        nil,
	//		PlayUrl:       utils.VideoUrlPrefix + v.PlayUrl,
	//		CoverUrl:      utils.ImageUrlPrefix + v.CoverUrl,
	//		FavoriteCount: v.FavoriteCount,
	//		CommentCount:  v.CommentCount,
	//		IsFavorite:    false,
	//	}
	//	res = append(res, video)
	//}
	//return res, nil
}

// GetVideoFeed 获取视频流，并查询follow及favorite更新is_follow字段
func (VideoService) GetVideoFeed(userId, t int64) ([]vo.Video, error) {

	videos, err := videoDao.GetVideoListByTime(t)
	if err != nil {
		return nil, err
	}

	res, err := handleFollowCondition(videos, userId)
	if err != nil {
		return nil, err
	}
	return res, nil

	//toUserIdList := make([]int64, len(videos))
	//toVideoIdList := make([]int64, len(videos))
	//for i, v := range videos {
	//	toUserIdList[i] = v.AuthorId
	//	toVideoIdList[i] = v.ID
	//}
	//
	//// 查询视频的发布者是否是用户的关注
	//// 如果没有登录，则全部没关注
	//// 否则，通过查询数据库确定
	//followUserMap := make(map[int64]bool, len(toUserIdList))
	//followVideoMap := make(map[int64]bool, len(toVideoIdList))
	//// 没有登录
	//if userId == -1 {
	//	// 全部赋值为false
	//	for _, uId := range toUserIdList {
	//		followUserMap[uId] = false
	//		followVideoMap[uId] = false
	//	}
	//} else {
	//	wg := sync.WaitGroup{}
	//	wg.Add(2)
	//
	//	var err1, err2 error
	//	// 查询用户是否关注
	//	go func() {
	//		defer wg.Done()
	//		followUserMap, err1 = followDao.GetFollowUserIdByUserId(userId, toUserIdList)
	//	}()
	//
	//	// 查询视频是否点赞
	//	go func() {
	//		defer wg.Done()
	//		followVideoMap, err2 = likeDao.GetFollowVideoIdByUserId(userId, toVideoIdList)
	//	}()
	//	wg.Wait()
	//	if err1 != nil {
	//		return nil, err1
	//	}
	//	if err2 != nil {
	//		return nil, err2
	//	}
	//}
	//
	//videoList := make([]vo.Video, len(videos))
	//
	//for i, video := range videos {
	//	isFollowAuthor, ok := followUserMap[video.AuthorId]
	//	isFollowVideo, ok2 := followVideoMap[video.ID]
	//	if !ok {
	//		isFollowAuthor = false
	//	}
	//	if !ok2 {
	//		isFollowVideo = false
	//	}
	//	videoList[i] = vo.Video{Id: video.ID,
	//		Author: &vo.Author{
	//			Id:            video.AuthorId,
	//			Name:          video.Author.Name,
	//			FollowCount:   video.Author.FollowCount,
	//			FollowerCount: video.Author.FollowerCount,
	//			IsFollow:      isFollowAuthor,
	//		},
	//		PlayUrl:       utils.VideoUrlPrefix + video.PlayUrl,
	//		CoverUrl:      utils.ImageUrlPrefix + video.CoverUrl,
	//		FavoriteCount: video.FavoriteCount,
	//		CommentCount:  video.CommentCount,
	//		IsFavorite:    isFollowVideo,
	//		Title:         video.Title,
	//		CreatedAt:     video.CreatedAt,
	//	}
	//
	//}
	//
	//return videoList, nil
}

func handleFollowCondition(videos []model.Video, userId int64) ([]vo.Video, error) {
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
	// 没有登录或是自己的
	isLogin := userId != -1
	if !isLogin {
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

func (VideoService) HandleFollowCondition(videos []model.Video, userId int64) ([]vo.Video, error) {
	return handleFollowCondition(videos, userId)
}
