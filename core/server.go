package core

import (
	"bytedance-douyin/global"
	"bytedance-douyin/initialize"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
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

	// 引入pprof
	initPprof()

	Router := initialize.Routers()
	//pprof.Register(Router, "dev/pprof")

	Router.Static("/form-generator", "./resource/page")

	Router.StaticFS("/videos", http.Dir(global.GVA_CONFIG.File.VideoOutput))
	Router.StaticFS("/images", http.Dir(global.GVA_CONFIG.File.ImageOutput))

	// address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	port := global.GVA_CONFIG.Port
	ip := global.GVA_CONFIG.IP
	s := initServer("0.0.0.0:"+port, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", ip))

	global.GVA_LOG.Error(s.ListenAndServe().Error())

}
