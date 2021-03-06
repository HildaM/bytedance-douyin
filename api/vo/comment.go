package vo

import "bytedance-douyin/types"

// CommentActionRequest 评论、删除评论
type CommentActionRequest struct {
	UserId      int64  `form:"user_id"`
	VideoId     int64  `form:"video_id" binding:"required"`
	ActionType  int8   `form:"action_type" binding:"required,oneof=1 2"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

// CommentListRequest 评论列表
type CommentListRequest struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id" binding:"required"`
}

type Comment struct {
	Id         int64      `json:"id"`
	User       UserInfo   `json:"user"`
	Content    string     `json:"content"`
	CreateDate types.Time `json:"create_date" time_format:"06-02"`
}

type CommentResponseVo struct {
	CommentList []Comment `json:"comment_list" binding:"required"`
}

type CommentActionResponseVo struct {
	Comment Comment `json:"comment"`
}
