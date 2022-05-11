package repository

import (
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"fmt"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:37
 */
type UserDao struct{}

func (UserDao) GetUser(userId int) (model.UserDao, error) {
	user := model.UserDao{}
	if result := global.GVA_DB.Where("id = ?", userId).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (UserDao) RegisterUser(userBo bo.UserBo) (userId int64) {
	fmt.Println(userBo)
	user := model.UserDao{Name: userBo.Name, Password: userBo.Pwd}
	global.GVA_DB.Create(&user)
	userId = user.ID
	return
}
