package router

import (
	"bytedance-douyin/api"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description: 关注操作、关注列表、粉丝列表
 * @File: follow
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:34
 */
type FollowRouter struct{}

func (c *FollowRouter) InitFollowRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("relation")
	followRouter := baseRouter.Group("follow")
	followerRouter := baseRouter.Group("follower")
	followApi := api.GroupApp.FollowApi
	{
		// /relation/action
		baseRouter.POST("action/", followApi.Follow)
		// /relation/follow/list
		followRouter.GET("list/", followApi.FollowList)
	}
	{
		// /relation/follower/list
		followerRouter.GET("list/", followApi.FansList)
	}
}
