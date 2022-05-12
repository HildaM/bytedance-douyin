package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/service/bo"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: CommentService
 * @Version: 1.0.0
 * @Date: 2022/5/12 16:36
 */
type CommentService struct {}

func (CommentService) GetCommentList(userId, videoId int64) (*[]vo.Comment, error) {
	list, err := commentDao.GetCommentList(videoId)
	if err != nil {
		return nil, err
	}
	res := make([]vo.Comment, 0)
	followList, _ := followDao.GetToUserIdList(userId)
	followMap := make(map[int64]bool)
	for _, v := range followList {
		followMap[v] = true
	}
	for _, v := range *list {
		comment := vo.Comment{
			Id:         v.ID,
			User:       vo.UserInfo{
				Id:            v.User.Id,
				Name:          v.User.Name,
				FollowCount:   v.User.FollowCount,
				FollowerCount: v.User.FollowerCount,
				IsFollow:      followMap[v.User.Id],
			},
			Content:    v.Content,
			CreateDate: v.CreateDate,
		}
		res = append(res, comment)
	}
	return &res, nil
}

func (s CommentService) DeleteComment(commentDelete bo.CommentDelete) error {
	if err := commentDao.DeleteComment(commentDelete.CommentId); err != nil {
		return err
	}
	return nil
}

func (s CommentService) PostComment(post bo.CommentPost) error {
	if err := commentDao.PostComment(post); err != nil {
		return err
	}
	return nil
}