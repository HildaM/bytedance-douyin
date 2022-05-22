package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"bytedance-douyin/service/bo"
)

type LikeService struct{}

func (LikeService) GetLikeList(likeListInfo vo.FavoriteListVo) (vo.FavoriteResponseVo, error) {
	var favoriteVideoList vo.FavoriteResponseVo
	userId := likeListInfo.UserId
	myId := likeListInfo.MyId
	// 视频点赞列表
	videos, err := likeDao.GetLikeList(userId)

	if err != nil {
		return favoriteVideoList, err
	}

	videoList, err := GroupApp.VideoService.HandleFollowCondition(videos, myId)
	if err != nil {
		return favoriteVideoList, err
	}

	//videoList := make([]*vo.Video, 0, len(videos))
	//
	//// CoverUrl
	//for _, video := range videos {
	//	author, err := userDao.GetUser(video.AuthorId)
	//	if err != nil {
	//		return favoriteVideoList, err
	//	}
	//	// 获取is_follow
	//	isFollow := true
	//	count, err := followDao.GetFollowCount(bo.FollowBo{UserId: userId, ToUserId: video.AuthorId})
	//	if err != nil {
	//		return favoriteVideoList, err
	//	}
	//	if count == 0 {
	//		isFollow = false
	//	}
	//
	//	videoInfo := vo.Video{
	//		Id: video.ID,
	//		Author: &vo.Author{
	//			Id:            author.ID,
	//			Name:          author.Name,
	//			FollowCount:   author.FollowCount,
	//			FollowerCount: author.FollowerCount,
	//			IsFollow:      isFollow,
	//		},
	//		PlayUrl:       video.PlayUrl,
	//		CoverUrl:      video.CoverUrl,
	//		FavoriteCount: video.FavoriteCount,
	//		CommentCount:  video.CommentCount,
	//		IsFavorite:    true,
	//	}
	//	videoList = append(videoList, &videoInfo)
	//}
	favoriteVideoList.VideoList = videoList

	return favoriteVideoList, nil
}

func (LikeService) LikeOrCancel(likeInfo vo.FavoriteActionVo) (int8, error) {
	videoLikedBo := bo.VideoLikedBo{
		UserId:  likeInfo.UserId,
		VideoId: likeInfo.VideoId,
	}
	// 查询是否有该点赞记录
	count, err := likeDao.GetIsFavorite(videoLikedBo)
	action := likeInfo.ActionType
	if err != nil {
		return 0, err
	}
	switch {
	case count != 0 && action == 1:
		// 点过赞还点
		// 可不处理
		// FIXME: 可能之前取消点赞
	case count == 0 && action == 1:
		// 没点过赞点赞
		err = likeDao.LikeVideo(videoLikedBo)
	case count != 0 && action == 2:
		// 点过赞取消点赞
		err = likeDao.UnLikeVideo(videoLikedBo)
	}
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return 0, err
	}

	return action, nil
}
