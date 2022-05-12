package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: comment
 * @Version: 1.0.0
 * @Date: 2022/5/12 16:40
 */
type CommentDao struct{}

func (CommentDao) GetCommentList(videoId int64) (*[]bo.Comment, error) {
	comments := make([]bo.Comment, 0)
	result := global.GVA_DB.Preload("User").Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comments, nil
}

func (CommentDao) DeleteComment(CommentId int64) error {
	comment := model.Comment{
		Base: model.Base{
			ID: CommentId,
		},
	}
	if result := global.GVA_DB.Delete(comment); result.Error != nil {
		return result.Error
	}
	return nil
}

func (CommentDao) PostComment(post bo.CommentPost) error {
	comment := model.Comment{
		VideoId:     post.VideoId,
		UserId:      post.UserId,
		CommentText: post.CommentText,
	}
	if result := global.GVA_DB.Create(comment); result.Error != nil {
		return result.Error
	}
	return nil
}
