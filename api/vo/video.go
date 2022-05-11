package vo

// 点赞操作
type FavoriteActionVo struct {
	UserId     int64  `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	VideoId    int64  `form:"video_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required,oneof=1 2"`
}

// 点赞列表
type FavoriteListVo struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// 视频流
type FeedVo struct {
	LatestTime int32 `form:"latest_time"`
}

type VideoListVo struct {
	Token string `form:"token" binding:"required"`
}

type FeedResponseVo struct {
	NextTime  int64    `json:"next_time"`
	VideoList []*Video `json:"video_list" binding:"required"`
}

type PublishResponseVo struct {
	VideoList []*Video `json:"video_list" binding:"required"`
}

type FavoriteResponseVo struct {
	VideoList []*Video `json:"video_list" binding:"required"`
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
	Id            int64  `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	FollowCount   int64  `json:"follow_count" binding:"required"`
	FollowerCount int64  `json:"follower_count" binding:"required"`
	IsFollow      bool   `json:"is_follow" binding:"required"`
}
