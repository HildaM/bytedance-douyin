package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: response
 * @Version: 1.0.0
 * @Date: 2022/5/7 10:47
 */

const (
	ERROR   = 7
	SUCCESS = 0
)

type noData struct {}

func outPutString(code int, msg string, data interface{}) string {
	comma := ""
	// 判断data是否不为空数据，即返回携带数据项。
	if _, ok := data.(noData); !ok {
		comma = ","
	}
	bytes, _ := json.Marshal(data)
	status := fmt.Sprintf("%s\"status_code\":%d,\"status_msg\":\"%s\"}", comma, code, msg)
	sb := strings.Builder{}
	// 去掉data json最后一个 } 字符
	sb.Write(bytes[:len(bytes)-1])
	sb.WriteString(status)
	return sb.String()
}

func Result(code int, message string, data interface{}, c *gin.Context) {
	c.Writer.Header().Set("Content-Type","application/json; charset=utf-8")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteHeaderNow()
	c.Writer.WriteString(outPutString(code, message, data))
}

func Ok(c *gin.Context) {
	Result(SUCCESS, "操作成功", noData{}, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, noData{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, "操作成功", data, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, message, data, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "操作失败", noData{}, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, noData{}, c)
}

func FailWithData(data interface{}, c *gin.Context) {
	Result(ERROR, "操作成功", data, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, message, data, c)
}

