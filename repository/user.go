package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:37
 */
type UserDao struct{}

func (userDao *UserDao) GetUser(userId int) (model.User, error) {
	user := model.User{}
	if result := global.GVA_DB.Where("id = ?", userId).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
