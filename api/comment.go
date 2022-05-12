package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/service/bo"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
 * @Author: 1999single
 * @Description: 评论操作、评论列表
 * @File: comment
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:33
 */

const (
	POST = 1
	DELETE = 2
)

type CommentApi struct{}

func (api *CommentApi) CommentOPS(c *gin.Context) {
	request := vo.CommentActionRequest{}
	if err := c.ShouldBind(&request); err != nil {
		response.FailWithMessage(c, fmt.Sprintf("%s", err))
		return
	}
	var err error
	if request.ActionType == POST {
		commentPost := bo.CommentPost{UserId: request.UserId, VideoId: request.VideoId, CommentText: request.CommentText}
		err = commentService.PostComment(commentPost)
	} else if request.ActionType == DELETE {
		commentDelete := bo.CommentDelete{UserId: request.UserId, VideoId: request.VideoId, CommentId: request.CommentId}
		err = commentService.DeleteComment(commentDelete)
	}
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithMessage(c, "操作成功")
}

func (api *CommentApi) CommentList(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	list, err := commentService.GetCommentList(int64(videoId))
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	fmt.Println(list)
	response.OkWithData(c, bo.Data{CommentList: *list})
}
