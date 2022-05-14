package core

import (
	"bytedance-douyin/global"
	"bytedance-douyin/initialize"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
	"runtime"
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
	Router.Static("/form-generator", "./resource/page")

	Router.StaticFS("/videos", http.Dir(global.GVA_CONFIG.File.VideoOutput))
	Router.StaticFS("/images", http.Dir(global.GVA_CONFIG.File.ImageOutput))

	// address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer("0.0.0.0:8080", Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", "127.0.0.1"))

	global.GVA_LOG.Error(s.ListenAndServe().Error())

}

func initPprof() {
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	// 引入pprof
	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		global.GVA_LOG.Info("Starting pprof to monitor this server")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			global.GVA_LOG.Error(err.Error())
		}
	}()
}
