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
 * @Description: 赞操作、点赞列表
 * @File: like
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:35
 */
type LikeApi struct{}

// Like 点赞操作
func (api *LikeApi) Like(c *gin.Context) {
	var likeInfo vo.FavoriteActionVo
	if err := c.ShouldBind(&likeInfo); err != nil {
		r.FailWithMessage(c, exceptions.ParamValidationError.Error())
	}
	code, err := likeService.LikeOrCancel(likeInfo)
	if err != nil {
		r.FailWithMessage(c, err.Error())
		return
	}
	action := func(code int8) string {
		if code == 1 {
			return "点赞"
		}
		return "取消点赞"
	}(code)
	r.OkWithMessage(c, action+"成功")
}

func (api *LikeApi) LikeList(c *gin.Context) {
	var likeListInfo vo.FavoriteListVo
	if err := c.ShouldBind(&likeListInfo); err != nil {
		r.FailWithMessage(c, "参数校验失败")
	}
	likeList, err := likeService.GetLikeList(likeListInfo)
	if err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	r.OkWithData(c, likeList)

}
