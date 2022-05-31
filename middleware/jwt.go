package middleware

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/exceptions"
	"bytedance-douyin/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证请求是否携带token
		//token := ""
		token := c.Query("token")
		if c.FullPath() == utils.VideoPostPath {
			token = c.PostForm("token")
		}
		//fmt.Println(token)
		// todo 为了方便测试，token为fangkaiwo时通过身份验证。
		if token == "fangkaiwo" {
			c.Next()
			return
		}
		if token == "" {
			response.FailWithMessage(c, exceptions.RejectRequestError.Error())
			c.Abort()
			return
		}
		// 解析token
		j := utils.NewJWT()
		//claims, err := j.ParseToken(token)
		claims, err := j.ParseTokenRedis(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithMessage(c, exceptions.LoginExpired.Error())
				c.Abort()
				return
			}
			response.FailWithMessage(c, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims.BaseClaims)
		c.Set("tokenId", claims.BaseClaims.Id)
		c.Next()
	}
}
