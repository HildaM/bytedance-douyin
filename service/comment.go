package service

import "bytedance-douyin/service/bo"

/**
 * @Author: 1999single
 * @Description:
 * @File: CommentService
 * @Version: 1.0.0
 * @Date: 2022/5/12 16:36
 */
type CommentService struct {}

func (CommentService) GetCommentList(videoId int64) (*[]bo.Comment, error) {
	list, err := commentDao.GetCommentList(videoId)
	if err != nil {
		return nil, err
	}
	return list, nil
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