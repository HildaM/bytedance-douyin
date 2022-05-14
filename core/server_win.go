package core

import (
	"bytedance-douyin/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: server_win
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:26
 */
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
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
