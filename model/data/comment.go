package data

type CommentData struct {
	CommentList []*Comment `json:"comment_list" binding:"required"`
}

type Comment struct {
	Id         int64     `json:"id"`
	User       *UserInfo `json:"user"`
	Content    string    `json:"content"`
	CreateDate string    `json:"create_date" time_format:"01-02"`
}
