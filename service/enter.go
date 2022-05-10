package service

import "bytedance-douyin/repository"

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:17
 */
type Group struct {
	UserService UserService
}

var GroupApp = new(Group)

// repository
var (
	userDao = repository.GroupApp.UserDao
)
