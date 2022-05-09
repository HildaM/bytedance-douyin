package global

import (
	"bytedance-douyin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
 * @Author: 1999single
 * @Description: 全局资源
 * @File: global
 * @Version: 1.0.0
 * @Date: 2022/5/6 15:29
 */

var (
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)

// 数据库相关
var (
	MYSQL_URL string
)
