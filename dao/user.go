// Package dao
// @Description: User Dao接口
// @Author: Quan
package dao

import (
	"bytedance-douyin/initialize"
	"gorm.io/gorm"
	"log"
	"sync"
)

type User struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type UserDao struct{}

var userDao *UserDao
var userOnce sync.Once // 确保只创建一次

func NewUserDanInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

// 获取所有user
func (*UserDao) GetAllUser() (*[]User, error) {
	var users []User
	err := initialize.Mysql.Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &users, nil
}
