package response

type FeedResponse struct {
	*BasicResponse
	NextTime  int64    `json:"next_time"`
	VideoList []*Video `json:"video_list"`
}

type PublishResponse struct {
	*BasicResponse
	VideoList []*Video `json:"video_list"`
}

type FavoriteResponse struct {
	*PublishResponse
}

type Video struct {
	Id            int64   `json:"id"`
	Author        *Author `json:"author"`
	PlayUrl       string  `json:"play_url"`
	CoverUrl      string  `json:"cover_url"`
	FavoriteCount int64   `json:"favorite_count"`
	CommentCount  int64   `json:"comment_count"`
	IsFavorite    bool    `json:"is_favorite"`
}

type Author struct {
	*UserInfo
}
