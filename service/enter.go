package service

import "bytedance-douyin/service/backend"

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:17
 */
type ServiceGroup struct {
	BackendServiceGroup  backend.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)