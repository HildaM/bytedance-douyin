package backend

import (
	response2 "bytedance-douyin/model/data"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description: 用户登陆、注册、查看信息
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:54
 */
type UserApi struct{}

func (api *UserApi) Register(c *gin.Context) {

}

func (api *UserApi) Login(c *gin.Context) {
	user := response2.UserInfoData{User: &response2.UserInfo{Id: 1, Name: "123", FollowCount: 1, FollowerCount: 1, IsFollow: true}}
	response2.OkWithData(c, user)
}

func (api *UserApi) UserInfo(c *gin.Context) {

}
