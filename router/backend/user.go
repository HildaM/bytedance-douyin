package backend

import (
	"bytedance-douyin/api/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:34
 */

type UserRouter struct{}

func (u *UserRouter) InitBaseUserRouter(Router *gin.RouterGroup) {
	baseUserRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.BackendApiGroup.UserApi
	{
		baseUserRouter.POST("login", userApi.Login)
		baseUserRouter.POST("register/", userApi.Register)
	}
}

func (u *UserRouter) InitUserInfoRouter(Router *gin.RouterGroup) {
	baseUserRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.BackendApiGroup.UserApi
	{
		baseUserRouter.GET("", userApi.UserInfo)
	}
}
