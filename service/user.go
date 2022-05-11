package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
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
	token := userInfo.Token
	isMyself, toId, err := utils.DoubleCheckToken(userId, token)
	userInfoBo := bo.UserInfoBo{}
	if err != nil {
		return userInfoBo, err
	}

	var userModel model.UserDao
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

	userInfoBo.ID = userModel.ID
	userInfoBo.Name = userModel.Name
	//  TODO 相关接口待实现
	//userInfoBo.FollowCount = userModel.FollowCount
	//userInfoBo.FollowerCount = userModel.FollowerCount
	userInfoBo.Follow = false
	return userInfoBo, nil
}

func (UserService) RegisterUser(user vo.UserVo) (userId int64) {
	var userBo bo.UserBo
	userBo.Name = user.Username
	userBo.Pwd = encrypt.Md5([]byte(user.Password))
	userId = userDao.RegisterUser(userBo)
	return
}

func (UserService) LoginUser(user vo.UserVo) (userId int64) {
	var userBo bo.UserBo
	userBo.Name = user.Username
	userBo.Pwd = encrypt.Md5([]byte(user.Password))
	userId = userDao.LoginUser(userBo)
	return
}
