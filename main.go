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
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	// zap.ReplaceGlobals(global.GVA_LOG)

	initialize.Redis() // 初始化redis数据库
	initialize.GormMysql() // 初始化mysql

	core.RunWindowsServer()
}
