package model

type Like struct {
	Base
	UserId  int64 `gorm:"type:bigint"`
	VideoId int64 `gorm:"type:bigint"`
}

func (Like) TableName() string {
	return "t_like"
}
