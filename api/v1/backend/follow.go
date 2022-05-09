package backend

import (
	r "bytedance-douyin/model/data"
	"bytedance-douyin/model/request"
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
	var param request.FollowRequest
	if err := c.ShouldBind(&param); err != nil {
		r.FailWithMessage(c, "参数校验失败")
	}

}

func (api *FollowApi) FollowList(c *gin.Context) {

}

func (api *FollowApi) FansList(c *gin.Context) {

}
