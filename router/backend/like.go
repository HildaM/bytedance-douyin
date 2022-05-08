package backend

import (
	v1 "bytedance-douyin/api/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description: 赞操作、点赞列表
 * @File: like
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:35
 */
type LikeRouter struct {}

func (l *LikeRouter) InitLikeRouter(Router *gin.RouterGroup)  {
	router := Router.Group("favorite")
	likeApi := v1.ApiGroupApp.BackendApiGroup.LikeApi
	{
		router.POST("action", likeApi.Like)
		router.GET("list", likeApi.LikeList)
	}
}