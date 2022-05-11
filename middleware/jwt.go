package middleware

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/utils"
	"github.com/gin-gonic/gin"
)

const (
	REJECT_REQUEST = "拒绝访问"
	LOGIN_EXPIRED  = "登录状态已过期"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证请求是否携带token
		token := ""
		switch c.Request.Method {
		case "GET":
			token = c.Query("token")
		case "POST":
			token = c.PostForm("token")
		}
		if token == "" {
			response.FailWithMessage(c, REJECT_REQUEST)
			c.Abort()
			return
		}
		// 解析token
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithMessage(c, LOGIN_EXPIRED)
				c.Abort()
				return
			}
			response.FailWithMessage(c, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims.BaseClaims)
		c.Next()
	}
}
