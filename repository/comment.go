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

func (CommentDao) GetCommentList(videoId int64) ([]model.Comment, error) {
	comments := make([]model.Comment, 0)
	result := global.GVA_DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Table(model.UserDao{}.TableName())
	}).Where("video_id = ?", videoId).Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (CommentDao) DeleteComment(CommentId int64) error {
	comment := model.Comment{
		Base: model.Base{
			ID: CommentId,
		},
	}
	tx := global.GVA_DB.Begin()

	if result := tx.Debug().Delete(&comment); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	videoId := comment.VideoId
	err := GroupApp.VideoDao.VideoCommentCountIncr(videoId, -1)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
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
	tx := global.GVA_DB.Begin()

	if result := tx.Debug().Create(&comment); result.Error != nil {
		tx.Rollback()
		return comment, result.Error
	}

	err := GroupApp.VideoDao.VideoCommentCountIncr(post.VideoId, 1)
	if err != nil {
		tx.Rollback()
		return comment, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return comment, err
	}

	return comment, nil
}
