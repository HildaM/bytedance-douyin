package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/service/bo"
	"github.com/u2takey/go-utils/encrypt"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:18
 */
type UserService struct{}

func (userService UserService) GetUserInfo(userId int) (bo.UserInfoBo, error) {

	userInfoBo := bo.UserInfoBo{}
	userModel, err := userDao.GetUser(userId)
	if err != nil {
		return userInfoBo, err
	}
	userInfoBo.ID = userModel.ID
	userInfoBo.Name = userModel.Name
	//userInfoBo.FollowCount = userModel.FollowCount
	//userInfoBo.FollowerCount = userModel.FollowerCount
	// 相关接口待实现
	userInfoBo.Follow = false
	return userInfoBo, nil
}

func (UserService) RegisterUser(user vo.UserVo) {
	var userBo bo.UserBo
	userBo.Name = user.Username
	userBo.Pwd = encrypt.Md5([]byte(user.Password))
	userDao.RegisterUser(userBo)

}
