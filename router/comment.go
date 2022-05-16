package router

import (
	"bytedance-douyin/api"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description: 评论操作、评论列表
 * @File: comment
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:33
 */
type CommentRouter struct{}

func (c *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	router := Router.Group("comment")
	commentApi := api.GroupApp.CommentApi
	{
		router.POST("action/", commentApi.CommentOPS)
		router.GET("list", commentApi.CommentList)
	}
}
