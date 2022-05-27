package bo

import (
	"bytedance-douyin/types"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: comment
 * @Version: 1.0.0
 * @Date: 2022/5/12 17:05
 */
type CommentPost struct {
	UserId      int64
	VideoId     int64
	CommentText string
}

type CommentDelete struct {
	UserId    int64
	VideoId   int64
	CommentId int64
}

type Comment struct {
	ID         int64      `json:"id"`
	User       UserInfoBo `json:"user"`
	UserId     int64      `json:"user_id"`
	Content    string     `json:"content"`
	CreateDate types.Time `gorm:"created_at" json:"create_date" time_format:"06-02"`
}
