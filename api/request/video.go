package request

// 点赞操作
type FavoriteActionRequest struct {
	UserId     int64  `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	VideoId    int64  `form:"video_id" binding:"required"`
	ActionType int8   `form:"action_type" binding:"required,oneof=1 2"`
}

// 点赞列表
type FavoriteListRequest struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// 视频流
type FeedRequest struct {
	LatestTime int32 `form:"latest_time"`
}

type VideoListRequest struct {
	Token string `form:"token" binding:"required"`
}
