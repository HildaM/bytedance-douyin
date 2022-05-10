package model

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/10 23:30
 */
type User struct {
	Base
	Name string
	Password string
	FollowCount int
	FollowerCount int
}

// 表名
func (User) TableName() string {
	return "t_user"
}