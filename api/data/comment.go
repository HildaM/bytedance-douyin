package data

import "bytedance-douyin/api/vo"

type Comment struct {
	CommentList []*vo.Comment `json:"comment_list" binding:"required"`
}


