package config

/**
 * @Author: 1999single
 * @Description: 应用配置
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/5/6 16:01
 */

type Server struct {
	Zap   Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	File  File   `mapstructure:"file" json:"file" yaml:"file"`
	Redis Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT   JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	IP    string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Port  string `mapstructure:"port" json:"port" yaml:"port"`
}
