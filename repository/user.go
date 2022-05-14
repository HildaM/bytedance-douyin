package repository

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/exceptions"
	"bytedance-douyin/global"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
	"errors"
	"gorm.io/gorm"
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
func (UserDao) GetUserByName(name string) error {
	var user model.UserDao
	err := global.GVA_DB.Where("name = ?", name).First(&user).Error
	// 没有该条记录
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	} else if err != nil {
		return err
	}
	// 有该条记录
	return exceptions.UserExistError
}

func (UserDao) GetUsers(userIds []int64) ([]model.UserDao, error) {
	var users []model.UserDao
	if result := global.GVA_DB.Where("id in (?)", userIds).Find(&users); result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (u UserDao) RegisterUser(userBo bo.UserBo) (bo.UserRegisterBo, error) {
	tx := global.GVA_DB.Begin()
	user := model.UserDao{Name: userBo.Name, Password: userBo.Pwd}
	var urb bo.UserRegisterBo
	// 判断用户名是否已存在
	err := u.GetUserByName(userBo.Name)
	if err != nil {
		return urb, err
	}

	// 不存在，创建用户
	tx.Debug().Create(&user)
	if tx.Error != nil {
		tx.Rollback()
		return urb, tx.Error
	}
	userId := user.ID

	bc := vo.BaseClaims{Id: userId, Name: userBo.Name}
	// 生成token
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

func (UserDao) QueryUserByNameAndPassword(userBo bo.UserBo) (int64, error) {
	var user model.UserDao
	err := global.GVA_DB.
		Where("name = ? and password = ?", userBo.Name, userBo.Pwd).
		Select("id").
		Find(&user).Error
	if err != nil {
		return 0, err
	}
	if user.ID == 0 {
		return 0, exceptions.LoginError
	}
	return user.ID, nil
}
