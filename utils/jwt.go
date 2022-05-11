package utils

import (
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: jwt
 * @Version: 1.0.0
 * @Date: 2022/5/11 18:16
 */

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims vo.BaseClaims) vo.CustomClaims {
	return vo.CustomClaims{
		BaseClaims: baseClaims,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,                          // 签名的发行者
		},
	}
}

func (j *JWT) CreateToken(claims vo.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenStr string) (*vo.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &vo.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*vo.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func (j JWT) ParseTokenRedis(token string) (*vo.CustomClaims, error) {
	r, err := global.GVA_REDIS.Get(context.Background(), token).Result()
	if err == nil {
		var cc vo.CustomClaims
		err := jsoniter.Unmarshal([]byte(r), &cc)
		if err != nil {
			return nil, err
		}
		return &cc, nil
	}
	cc, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func GenerateAndSaveToken(claims vo.BaseClaims) (string, error) {
	myJwt := NewJWT()
	customClaim := myJwt.CreateClaims(claims)
	token, err := myJwt.CreateToken(customClaim)
	if err != nil {
		return "", err
	}

	et := global.GVA_CONFIG.JWT.ExpiresTime
	if err := global.GVA_REDIS.SetEX(context.Background(), token, &customClaim, time.Duration(et)*time.Second).Err(); err != nil {
		return "", err
	}
	return token, nil
}
