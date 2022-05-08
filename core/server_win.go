package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
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