package api

import (
	r "bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/exceptions"
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

// LikeList 点赞列表
func (api *LikeApi) LikeList(c *gin.Context) {
	var likeListInfo vo.FavoriteListVo
	if err := c.ShouldBind(&likeListInfo); err != nil {
		r.FailWithMessage(c, exceptions.ParamValidationError.Error())
	}

	tokenId, ok := c.Get("tokenId")
	if !ok {
		r.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}

	likeListInfo.MyId = tokenId.(int64)
	likeList, err := likeService.GetLikeList(likeListInfo)
	if err != nil {
		r.FailWithMessage(c, err.Error())
	}

	r.OkWithData(c, likeList)

}
