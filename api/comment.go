package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/exceptions"
	"bytedance-douyin/service/bo"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description: 评论操作、评论列表
 * @File: comment
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:33
 */

const (
	POST   = 1
	DELETE = 2
)

type CommentApi struct{}

// CommentOPS 添加评论或删除评论
func (api *CommentApi) CommentOPS(c *gin.Context) {
	request := vo.CommentActionRequest{}
	if err := c.ShouldBind(&request); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	tokenId, ok := c.Get("tokenId")
	if !ok {
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}
	request.UserId = tokenId.(int64)

	var err error
	var comment vo.Comment

	data := vo.CommentActionResponseVo{}

	switch request.ActionType {
	case POST:
		commentPost := bo.CommentPost{UserId: request.UserId, VideoId: request.VideoId, CommentText: request.CommentText}
		comment, err = commentService.PostComment(commentPost)
	case DELETE:
		commentDelete := bo.CommentDelete{UserId: request.UserId, VideoId: request.VideoId, CommentId: request.CommentId}
		err = commentService.DeleteComment(commentDelete)
	}

	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 添加评论返回评论信息
	if request.ActionType == POST {
		var info bo.UserInfoBo
		info, err = userService.GetUserInfo(vo.UserInfoVo{UserId: request.UserId, MyUserId: request.UserId})
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		comment.User = vo.UserInfo{Id: info.Id,
			Name:          info.Name,
			FollowCount:   info.FollowCount,
			FollowerCount: info.FollowerCount,
			IsFollow:      info.Follow,
		}
		data.Comment = comment
	}
	response.OkWithDetailed(c, "操作成功", data)
}

func (api *CommentApi) CommentList(c *gin.Context) {
	var request vo.CommentListRequest

	if err := c.ShouldBind(&request); err != nil {
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}

	list, err := commentService.GetCommentList(request.UserId, request.VideoId)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	response.OkWithData(c, vo.CommentResponseVo{CommentList: list})
}
