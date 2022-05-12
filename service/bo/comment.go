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

type Data struct {
	CommentList []Comment `json:"comment_list"`
}

type Comment struct {
	ID         int64 `json:"id"`
	User       User `json:"user"`
	UserId     int64 `json:"user_id"`
	Content    string `json:"content"`
	CreateDate types.Time `gorm:"column:created_at" json:"create_date"`
}

func (Comment) TableName() string {
	return "t_comment"
}

type User struct {
	ID            int64 `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
	Follow        bool `json:"is_follow"`
}

func (User) TableName() string {
	return "t_user"
}
