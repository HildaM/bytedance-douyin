package backend

import (
	"bytedance-douyin/repository"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:18
 */
type ServiceGroup struct {
	UserService UserService
}

// repository
var (
	userDao = repository.RepositoryGroupApp.BackendRepositoryGroup.UserDao
)