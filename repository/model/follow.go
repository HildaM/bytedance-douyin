// Package model
// @Description: t_follow实体
// @Author: Quan

package model

type Follow struct {
	Base
	UserId   int64 `gorm:"type:bigint"`
	ToUserId int64 `gorm:"type:bigint"`
}

func (Follow) TableName() string {
	return "t_follow"
}
