package router

import "bytedance-douyin/router/backend"

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:27
 */

type RouterGroup struct {
	Backend backend.RouterGroup

}

var RouterGroupApp = new(RouterGroup)