package model

/**
 * @Author: 1999single
 * @Description:
 * @File: comment
 * @Version: 1.0.0
 * @Date: 2022/5/12 16:31
 */
type Comment struct {
	Base
	VideoId int64
	UserId  int64
	Content string
}

// 表名
func (Comment) TableName() string {
	return "t_comment"
}
