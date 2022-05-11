package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:37
 */
type UserDao struct{}

func (UserDao) GetUser(userId int64) (model.UserDao, error) {
	user := model.UserDao{}
	if result := global.GVA_DB.Where("id = ?", userId).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (UserDao) RegisterUser(userBo bo.UserBo) (userId int64) {
	user := model.UserDao{Name: userBo.Name, Password: userBo.Pwd}
	global.GVA_DB.Create(&user)
	userId = user.ID
	return
}

func (u *UserDao) LoginUser(userBo bo.UserBo) (userId int64) {
	var userLoginBo bo.UserLoginBo
	global.GVA_DB.Where("username = ? and password = ?", userBo.Name, userBo.Pwd).Find(&userLoginBo)
	return userLoginBo.Id
}
