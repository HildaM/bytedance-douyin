package initialize

import (
	"bytedance-douyin/global"
	"bytedance-douyin/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() {
	m := global.GVA_CONFIG.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		SkipInitializeWithVersion: false,
	}
	if gormMysql, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err == nil {
		db, _ := gormMysql.DB()
		db.SetMaxIdleConns(m.MaxIdleConns)
		db.SetMaxOpenConns(m.MaxOpenConns)
		global.GVA_DB = gormMysql
	} else {
		global.GVA_LOG.Error(err.Error())
	}
}
