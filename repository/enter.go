package repository

import "bytedance-douyin/repository/backend"

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:26
 */

type RepositoryGroup struct {
	BackendRepositoryGroup backend.RepositoryGroup
}

var RepositoryGroupApp = new(RepositoryGroup)