package vo

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 1:27
 */
type UserInfo struct {
	Id            int64  `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	FollowCount   int  `json:"follow_count" binding:"required"`
	FollowerCount int  `json:"follower_count" binding:"required"`
	Follow      bool   `json:"is_follow" binding:"required"`
}