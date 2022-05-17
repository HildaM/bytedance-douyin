package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/exceptions"
	"bytedance-douyin/utils"
	"github.com/gin-gonic/gin"
)

// UserApi
/**
用户注册、登录、获取用户信息接口
@author: charon
@date:   2022-5-14 last update
*/
type UserApi struct{}

func (UserApi) Register(c *gin.Context) {
	var userRegister vo.UserVo
	if err := c.ShouldBind(&userRegister); err != nil {
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
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
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}
	userId, err := userService.LoginUser(userLogin)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
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
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}

	// 一定要拿到，否则panic
	claim := c.MustGet("claims")

	claims := claim.(vo.BaseClaims)
	userInfo.MyUserId = claims.Id
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
