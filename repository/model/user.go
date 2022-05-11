package model

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/10 23:30
 */
type UserDao struct {
	Base
	Name          string `gorm:"type:varchar(32)"`
	Password      string `gorm:"type:varchar(32)"`
	FollowCount   int64
	FollowerCount int64
}

// 表名
func (UserDao) TableName() string {
	return "t_user"
}
