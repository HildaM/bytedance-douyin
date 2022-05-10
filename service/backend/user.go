package backend

import "bytedance-douyin/repository/model"

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:18
 */
type UserService struct {}

func (userService *UserService) GetUserInfo(userId int) (model.User, error) {


	userModel, err := userDao.GetUser(userId);
	if err != nil {
		return userModel, err
	}
	return userModel, nil
}