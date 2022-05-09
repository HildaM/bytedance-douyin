package backend

import (
	v1 "bytedance-douyin/api/v1"
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
	followApi := v1.ApiGroupApp.BackendApiGroup.FollowApi
	{
		// /relation/follow/action
		followRouter.POST("action", followApi.Follow)
		// /relation/follow/list
		followRouter.GET("list", followApi.FansList)
	}
	{
		// /relation/follower/list
		followerRouter.GET("list", followApi.FollowList)
	}
}
