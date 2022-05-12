package repository

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
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

func (UserDao) GetUsers(userIds []int64) ([]model.UserDao, error) {
	var users []model.UserDao
	if result := global.GVA_DB.Where("id in (?)", userIds).Find(&users); result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (UserDao) RegisterUser(userBo bo.UserBo) (bo.UserRegisterBo, error) {
	tx := global.GVA_DB.Begin()
	user := model.UserDao{Name: userBo.Name, Password: userBo.Pwd}
	var urb bo.UserRegisterBo
	tx.Debug().Create(&user)
	if tx.Error != nil {
		tx.Rollback()
		return urb, tx.Error
	}
	userId := user.ID

	bc := vo.BaseClaims{Id: userId, Name: userBo.Name}
	token, err := utils.GenerateAndSaveToken(bc)
	if err != nil {
		tx.Rollback()
		return urb, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return urb, err
	}
	urb.Token = token
	urb.Id = userId
	return urb, nil
}

func (UserDao) QueryUserByNameAndPassword(userBo bo.UserBo) (userId int64) {
	var user model.UserDao
	global.GVA_DB.Where("name = ? and password = ?", userBo.Name, userBo.Pwd).Select("id").Find(&user)
	return user.ID
}
