package repository

import (
	"bytedance-douyin/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestMain(m *testing.M) {
	dsn := "root:220429douyin@tcp(220.243.147.162:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.GVA_DB = db
}

func TestFollowDao_GetFollowList(t *testing.T) {
	var followDao FollowDao
	list, err := followDao.GetFollowList(1)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	for _, user := range list.UserList {
		fmt.Println(user)
	}
}
