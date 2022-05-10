package initialize

import (
	"bytedance-douyin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

// 初始化数据库连接，dsn为具体的数据库连接地址
func Init() error {
	var err error
	dsn := global.MYSQL_URL
	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
