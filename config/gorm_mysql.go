package config

/**
 * @Author: 1999single
 * @Description:
 * @File: mysql
 * @Version: 1.0.0
 * @Date: 2022/5/10 21:55
 */
type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // 端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}