// Package repository
// @Description: Video Dao接口
// @Author: Quan

package repository

import (
	"bytedance-douyin/api/data"
)

type Video struct {
	Id            int64       `gorm:"column:id"`
	Author        data.Author `gorm:"column:author"`
	PlayUrl       string      `gorm:"column:play_url"`
	CoverUrl      string      `gorm:"column:cover_url"`
	FavoriteCount int         `gorm:"column:favorite_url"`
	CommentCount  int         `gorm:"column:comment_count"`
	IsFavorite    bool        `gorm:"is_favorite"`
}
