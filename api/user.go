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

func (api *UserApi) Register(c *gin.Context) {
	var userRegister vo.UserVo
	if err := c.ShouldBind(&userRegister); err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	userId := userService.RegisterUser(userRegister)
	bc := vo.BaseClaims{Id: userId, Name: userRegister.Username}
	token, err := utils.GenerateAndSaveToken(bc)
	if err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	urv := vo.UserResponseVo{UserId: userId, Token: token}
	response.OkWithData(c, urv)
}

func (api *UserApi) Login(c *gin.Context) {
	var userLogin vo.UserVo
	if err := c.ShouldBind(&userLogin); err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	userId := userService.LoginUser(userLogin)
	// BUG: userId在用户不存在的时候没有返回

	bc := vo.BaseClaims{Id: userId, Name: userLogin.Username}
	token, err := utils.GenerateAndSaveToken(bc)
	if err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	urv := vo.UserResponseVo{UserId: userId, Token: token}
	response.OkWithData(c, urv)
}

func (api *UserApi) UserInfo(c *gin.Context) {
	var userInfo vo.UserInfoVo
	if err := c.ShouldBind(&userInfo); err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	//token := userInfo.Token
	userService.GetUserInfo(userInfo)
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
