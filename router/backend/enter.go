package backend

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:28
 */

type RouterGroup struct {
	UserRouter
	LikeRouter
	CommentRouter
	VideoRouter
	FollowRouter
}