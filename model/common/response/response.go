package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: Charon
 * @Description:
 * @File: response
 * @Version: 1.0.0
 * @Date: 2022/5/9 12:52
 */

const (
	ERROR           = 7
	SUCCESS         = 0
	SUCCESS_MESSAGE = "success"
	ERROR_MESSAGE   = "error"
)

type BasicResponse struct {
	StatusCode    int8   `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	_             interface{}
}

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: SUCCESS, StatusMessage: SUCCESS_MESSAGE})
}

func OkWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: SUCCESS, StatusMessage: message})
}

func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &data)
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: ERROR, StatusMessage: ERROR_MESSAGE})
}

func FailWithMessage(message string, c *gin.Context) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: ERROR, StatusMessage: message})
}

func FailWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &data)
}
