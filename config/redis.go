package config

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Db       int    `mapstructure:"db" json:"db" yaml:"db"`
	PoolSize int    `mapstructure:"poolsize" json:"poolsize" yaml:"poolsize"`
}
