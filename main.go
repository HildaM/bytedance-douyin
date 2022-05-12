package main

import (
	"bytedance-douyin/core"
	"bytedance-douyin/global"
	"bytedance-douyin/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	// zap.ReplaceGlobals(global.GVA_LOG)

	initialize.Redis()     // 初始化redis数据库
	initialize.GormMysql() // 初始化mysql

	//TestFollowDao_GetFollowList()
	core.RunWindowsServer()

}

//func TestFollowDao_GetFollowList() {
//	var followDao repository.FollowDao
//	list, err := followDao.GetFollowList(1)
//	if err != nil {
//		fmt.Errorf(err.Error())
//	}
//
//	for _, user := range list.UserList {
//		fmt.Println(user)
//	}
//}
