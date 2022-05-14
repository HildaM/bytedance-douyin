package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/utils"
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

func (UserApi) Register(c *gin.Context) {
	var userRegister vo.UserVo
	if err := c.ShouldBind(&userRegister); err != nil {
		response.FailWithMessage(c, err.Error())
	}

	urb, err := userService.RegisterUser(userRegister)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	urv := vo.UserResponseVo{UserId: urb.Id, Token: urb.Token}
	response.OkWithData(c, urv)
}

func (UserApi) Login(c *gin.Context) {
	var userLogin vo.UserVo
	if err := c.ShouldBind(&userLogin); err != nil {
		response.FailWithMessage(c, err.Error())
	}
	userId, err := userService.LoginUser(userLogin)
	if err != nil {
		response.FailWithMessage(c, err.Error())
	}

	bc := vo.BaseClaims{Id: userId, Name: userLogin.Username}
	token, err := utils.GenerateAndSaveToken(bc)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	urv := vo.UserResponseVo{UserId: userId, Token: token}
	response.OkWithData(c, urv)
}

func (api *UserApi) UserInfo(c *gin.Context) {
	var userInfo vo.UserInfoVo
	if err := c.ShouldBind(&userInfo); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	//token := userInfo.Token
	info, err := userService.GetUserInfo(userInfo)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	u := vo.UserInfo{Id: info.Id, Name: info.Name, FollowCount: info.FollowCount, FollowerCount: info.FollowerCount, IsFollow: info.Follow}

	data := vo.UserInfoResponseVo{User: u}
	response.OkWithData(c, data)

}
