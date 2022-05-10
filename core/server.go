package core

import (
	"bytedance-douyin/global"
	"bytedance-douyin/initialize"
	"go.uber.org/zap"
	"time"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:25
 */
type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	// address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer("0.0.0.0:8080", Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", "127.0.0.1"))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
