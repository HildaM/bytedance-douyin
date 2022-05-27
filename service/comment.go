package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/types"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: CommentService
 * @Version: 1.0.0
 * @Date: 2022/5/12 16:36
 */
type CommentService struct{}

// GetCommentList 获取评论列表
func (CommentService) GetCommentList(userId, videoId int64) ([]vo.Comment, error) {
	list, err := commentDao.GetCommentList(videoId)
	if err != nil {
		return nil, err
	}

	// 获取评论用户的id
	followList, _ := followDao.GetToUserIdListByRedis(userId)
	followMap := make(map[int64]bool, len(followList))
	for _, v := range followList {
		followMap[v] = true
	}

	res := make([]vo.Comment, len(list))
	for i, v := range list {
		comment := vo.Comment{
			Id: v.ID,
			User: vo.UserInfo{
				Id:            v.User.ID,
				Name:          v.User.Name,
				FollowCount:   v.User.FollowCount,
				FollowerCount: v.User.FollowerCount,
				IsFollow:      followMap[v.User.ID],
			},
			Content:    v.Content,
			CreateDate: types.Time(v.CreatedAt),
		}
		res[i] = comment
	}
	return res, nil
}

// DeleteComment 删除评论
func (s CommentService) DeleteComment(commentDelete bo.CommentDelete) error {
	if err := commentDao.DeleteComment(commentDelete.CommentId); err != nil {
		return err
	}
	return nil
}

// PostComment 添加评论
func (s CommentService) PostComment(post bo.CommentPost) (vo.Comment, error) {
	var comment vo.Comment
	commentPosted, err := commentDao.PostComment(post)
	if err != nil {
		return comment, err
	}
	comment = vo.Comment{
		Id:         commentPosted.ID,
		Content:    commentPosted.Content,
		CreateDate: types.Time(commentPosted.CreatedAt),
	}
	return comment, nil
}
