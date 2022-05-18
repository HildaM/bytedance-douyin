package service

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/repository/model"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
	"sync"
)

// UserService
/**
获取用户信息，登录、注册 service层
@author: charon
@date:   2022-5-14 last update
*/
type UserService struct{}

// GetUserInfo get user information
func (UserService) GetUserInfo(userInfo vo.UserInfoVo) (bo.UserInfoBo, error) {
	userId := userInfo.UserId
	myId := userInfo.MyUserId

	// judge whether me
	isMyself := userId == myId

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

		// use follow service judge whether I followed him
		var count int64
		count, err = GroupApp.FollowService.GetFollowCount(vo.FollowVo{UserId: myId, ToUserId: userId})
		if count != 0 {
			userInfoBo.Follow = true
		}
	}()

	wg.Wait()

	if err != nil {
		return userInfoBo, err
	}

	userInfoBo.Id = userModel.ID
	userInfoBo.Name = userModel.Name
	userInfoBo.FollowCount = userModel.FollowCount
	userInfoBo.FollowerCount = userModel.FollowerCount
	//  TODO 相关接口待实现
	//userInfoBo.FollowCount = userModel.FollowCount
	//userInfoBo.FollowerCount = userModel.FollowerCount
	return userInfoBo, nil
}

func (UserService) RegisterUser(user vo.UserVo) (bo.UserRegisterBo, error) {
	var userBo bo.UserBo
	userBo.Name = user.Username
	userBo.Pwd = utils.GetSHA256(user.Password)
	urb, err := userDao.RegisterUser(userBo)
	if err != nil {
		return urb, err
	}
	return urb, nil
}

func (UserService) LoginUser(user vo.UserVo) (int64, error) {
	var userBo bo.UserBo
	userBo.Name = user.Username
	userBo.Pwd = utils.GetSHA256(user.Password)
	userId, err := userDao.QueryUserByNameAndPassword(userBo)
	if err != nil {
		return userId, err
	}
	return userId, nil
}
