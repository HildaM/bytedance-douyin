package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"fmt"
	"github.com/u2takey/go-utils/encrypt"
	"sync"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:18
 */
type UserService struct{}

func (UserService) GetUserInfo(userInfo vo.UserInfoVo) (bo.UserInfoBo, error) {
	userId := userInfo.UserId
	toId := userInfo.Claims.BaseClaims.Id

	isMyself := userId == toId

	userInfoBo := bo.UserInfoBo{}

	var userModel model.UserDao
	var err error
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		userModel, err = userDao.GetUser(userId)
	}()

	go func() {
		defer wg.Done()
		if isMyself {
			userInfoBo.Follow = false
			return
		}
		// todo 查询是否是粉丝
		//FollowS
		fmt.Println(toId)
	}()

	wg.Wait()

	if err != nil {
		return userInfoBo, err
	}

	userInfoBo.Id = userModel.ID
	userInfoBo.Name = userModel.Name
	userInfoBo.FollowCount = 0
	userInfoBo.FollowerCount = 0
	// 相关接口待实现
	userInfoBo.Follow = false
	return userInfoBo, nil
}

func (UserService) RegisterUser(user vo.UserVo) (bo.UserRegisterBo, error) {
	var userBo bo.UserBo
	userBo.Name = user.Username
	userBo.Pwd = encrypt.Md5([]byte(user.Password))
	urb, err := userDao.RegisterUser(userBo)
	if err != nil {
		return urb, err
	}
	return urb, nil
}

func (UserService) LoginUser(user vo.UserVo) (userId int64) {
	var userBo bo.UserBo
	userBo.Name = user.Username
	userBo.Pwd = encrypt.Md5([]byte(user.Password))
	userId = userDao.LoginUser(userBo)
	return
}
