package backend

import (
	response "bytedance-douyin/model/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	user := response.UserInfoData{User: &response.UserInfo{Id: 1, Name: "123", FollowCount: 1, FollowerCount: 1, IsFollow: true}}
	response.OkWithData(c, user)
}

func (api *UserApi) UserInfo(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	user, err := userService.GetUserInfo(userId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
