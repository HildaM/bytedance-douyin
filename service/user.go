package service

import (
	"bytedance-douyin/service/bo"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:18
 */
type UserService struct{}

func (userService *UserService) GetUserInfo(userId int) (bo.User, error) {

	userBo := bo.User{}
	userModel, err := userDao.GetUser(userId)
	if err != nil {
		return userBo, err
	}
	userBo.ID = userModel.ID
	userBo.Name = userModel.Name
	userBo.FollowCount = userModel.FollowCount
	userBo.FollowerCount = userModel.FollowerCount
	// 相关接口待实现
	userBo.Follow = false
	return userBo, nil
}
