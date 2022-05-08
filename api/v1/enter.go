package v1

import "bytedance-douyin/api/v1/backend"

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:53
 */

type ApiGroup struct {
	BackendApiGroup   backend.ApiGroup
}

var ApiGroupApp = new(ApiGroup)