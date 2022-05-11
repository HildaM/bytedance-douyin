package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/utils"
	"fmt"
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
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	urb, err := userService.RegisterUser(userRegister)
	if err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
		return
	}

	urv := vo.UserResponseVo{UserId: urb.Id, Token: urb.Token}
	response.OkWithData(c, urv)
}

func (UserApi) Login(c *gin.Context) {
	var userLogin vo.UserVo
	if err := c.ShouldBind(&userLogin); err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	userId := userService.LoginUser(userLogin)
	bc := vo.BaseClaims{Id: userId, Name: userLogin.Username}
	token, err := utils.GenerateAndSaveToken(bc)
	if err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	urv := vo.UserResponseVo{UserId: userId, Token: token}
	fmt.Println(token)
	response.OkWithData(c, urv)
}

func (api *UserApi) UserInfo(c *gin.Context) {
	var userInfo vo.UserInfoVo
	if err := c.ShouldBind(&userInfo); err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
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
	//userId, _ := strconv.Atoi(c.Query("user_id"))
	//userBO, _ := userService.GetUserInfo(userId)
	//userVO := &vo.UserInfo{}
	//userVO.Id = userBO.ID
	//userVO.Name = userBO.Name
	//userVO.FollowerCount = userBO.FollowerCount
	//userVO.FollowCount = userBO.FollowCount
	//userVO.Follow = userBO.Follow
	//response.OkWithData(c, data.UserInfo{User: userVO})
}
