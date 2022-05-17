package api

import (
	r "bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/exceptions"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description: 关注操作、关注列表、粉丝列表
 * @File: follow
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:34
 */
type FollowApi struct{}

func (api *FollowApi) Follow(c *gin.Context) {
	var followInfo vo.FollowVo
	if err := c.ShouldBind(&followInfo); err != nil {
		r.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}

	var err error
	var code int8
	if code, err = followService.FollowOrNot(followInfo); err != nil {
		r.FailWithMessage(c, err.Error())
		return
	}
	action := func(code int8) string {
		if code == 1 {
			return "关注"
		}
		return "取消关注"
	}(code)

	r.OkWithMessage(c, action+"成功")
}

//
//  FollowList
//  @Description:	获取关注列表
//  @receiver api
//  @param c
//	@Request body:	user_id, token
// 	@Author Quan
//
func (api *FollowApi) FollowList(c *gin.Context) {
	var userInfo vo.FollowListVo
	if err := c.ShouldBind(&userInfo); err != nil {
		r.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}

	userList, err := followService.GetFollowList(userInfo)
	if err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
		return
	}
	r.OkWithData(c, userList)
}

//
//  FansList
//  @Description: 	获取粉丝列表
//  @receiver api
//	@Request body:	user_id, token
//  @param c
//
func (api *FollowApi) FansList(c *gin.Context) {
	var userInfo vo.FollowerListVo
	if err := c.ShouldBind(&userInfo); err != nil {
		r.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}
	tokenId, ok := c.Get("tokenId")
	if !ok {
		r.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}
	userInfo.TokenId = tokenId.(int64)

	fanList, err := followerService.GetFanList(userInfo)
	if err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
		return
	}

	r.OkWithData(c, fanList)
}
