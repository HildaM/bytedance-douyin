package router

import (
	"bytedance-douyin/api"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: vedio
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:33
 */
type VideoRouter struct{}

func (v *VideoRouter) InitVideoFeedRouter(Router *gin.RouterGroup) {
	router := Router.Group("feed")
	videoApi := api.GroupApp.VideoApi
	{
		router.GET("", videoApi.VideoFeed)
	}
}

func (v *VideoRouter) InitVideoRouter(Router *gin.RouterGroup) {
	router := Router.Group("publish")
	videoApi := api.GroupApp.VideoApi
	{
		router.POST("action/", videoApi.PostVideo)
		router.GET("list", videoApi.VideoList)
	}
}
