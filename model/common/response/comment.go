package response

type CommentResponse struct {
	*BasicResponse
	CommentList []*Comment
}

type Comment struct {
	Id         int64     `json:"id"`
	User       *UserInfo `json:"user"`
	Content    string    `json:"content"`
	CreateDate string    `json:"create_date"`
}
