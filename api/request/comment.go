package request

// 评论、删除评论
type CommentActionRequest struct {
	UserId      int64  `form:"user_id" binding:"required"`
	Token       string `form:"token" binding:"required"`
	VideoId     int64  `form:"video_id" binding:"required"`
	ActionType  int8   `form:"action_type" binding:"required,oneof=1 2"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

// 评论列表
type CommentListRequest struct {
	UserId  int64  `form:"user_id" binding:"required"`
	Token   string `form:"token" binding:"required"`
	VideoId int64  `form:"video_id" binding:"required"`
}
