package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"gorm.io/gorm"
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
	result := global.GVA_DB.Model(&model.Comment{}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Table(model.UserDao{}.TableName())
	}).Where("video_id = ?", videoId).Find(&comments)

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
	if result := global.GVA_DB.Delete(&comment); result.Error != nil {
		return result.Error
	}
	return nil
}

func (CommentDao) PostComment(post bo.CommentPost) (model.Comment, error) {
	var comment model.Comment
	comment = model.Comment{
		VideoId: post.VideoId,
		UserId:  post.UserId,
		Content: post.CommentText,
	}
	if result := global.GVA_DB.Create(&comment); result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}
