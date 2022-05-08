package backend

import "github.com/gin-gonic/gin"

/**
 * @Author: 1999single
 * @Description: 用户登陆、注册、查看信息
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/6 17:54
 */
type UserApi struct {}


func (api *UserApi) Register(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg": "success",
	})
}

func (api *UserApi) Login(c *gin.Context)  {

}

func (api *UserApi) UserInfo(c *gin.Context)  {

}