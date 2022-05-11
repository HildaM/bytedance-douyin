package config

/**
 * @Author: 1999single
 * @Description:
 * @File: jwt
 * @Version: 1.0.0
 * @Date: 2022/5/11 18:18
 */
type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}