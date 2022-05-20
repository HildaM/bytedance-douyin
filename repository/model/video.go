package model

type Video struct {
	Base
	AuthorId int64 `gorm:"type:bigint"`
	// VideoId       int64  `gorm:"type:bigint"`
	PlayUrl       string `gorm:"type:varchar(64)"`
	CoverUrl      string `gorm:"type:varchar(64)"`
	FavoriteCount int64  `gorm:"type:bigint"`
	CommentCount  int64  `gorm:"type:bigint"`
}

func (Video) TableName() string {
	return "t_video"
}
