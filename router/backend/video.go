package backend

import (
	v1 "bytedance-douyin/api/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: vedio
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:33
 */
type VideoRouter struct {}

func (v *VideoRouter) InitVideoFeedRouter(Router *gin.RouterGroup)  {
	router := Router.Group("feed")
	videoApi := v1.ApiGroupApp.BackendApiGroup.VideoApi
	{
		router.GET("", videoApi.VideoFeed)
	}
}

func (v *VideoRouter) InitVideoRouter(Router *gin.RouterGroup)  {
	router := Router.Group("publish")
	videoApi := v1.ApiGroupApp.BackendApiGroup.VideoApi
	{
		router.POST("action", videoApi.PostVideo)
		router.GET("list", videoApi.VideoList)
	}
}