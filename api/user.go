package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"fmt"
	"github.com/gin-gonic/gin"
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
	var userRegister vo.UserVo
	if err := c.ShouldBind(&userRegister); err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	userService.RegisterUser(userRegister)

}

func (api *UserApi) Login(c *gin.Context) {
	// 示例
	//user := data.UserInfo{User: &vo.UserInfo{Id: 1, Name: "123", FollowCount: 1, FollowerCount: 1, Follow: true}}
	//response.OkWithData(c, user)
}

func (api *UserApi) UserInfo(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	userBO, _ := userService.GetUserInfo(userId)
	userVO := &vo.UserInfo{}
	userVO.Id = userBO.ID
	userVO.Name = userBO.Name
	//userVO.FollowerCount = userBO.FollowerCount
	//userVO.FollowCount = userBO.FollowCount
	//userVO.Follow = userBO.Follow
	//response.OkWithData(c, data.UserInfo{User: userVO})
}
