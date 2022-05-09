package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
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
	Data          interface{}
	//_             interface{}
}

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: SUCCESS, StatusMessage: SUCCESS_MESSAGE})
}

func OkWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: SUCCESS, StatusMessage: message})
}

func OkWithData(c *gin.Context, data interface{}) {
	m, err := struct2Map(&BasicResponse{StatusCode: 0, StatusMessage: "success", Data: data})

	if err != nil {
		_ = fmt.Errorf("%w", err)
	}

	c.JSON(http.StatusOK, m)
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: ERROR, StatusMessage: ERROR_MESSAGE})
}

func FailWithMessage(message string, c *gin.Context) {
	c.JSON(http.StatusOK, &BasicResponse{StatusCode: ERROR, StatusMessage: message})
}

func FailWithData(c *gin.Context, data interface{}) {
	m, err := struct2Map(&BasicResponse{StatusCode: 0, StatusMessage: "success", Data: data})

	if err != nil {
		_ = fmt.Errorf("%w", err)
	}
	c.JSON(http.StatusOK, m)
}

func struct2Map(in interface{}) (map[string]interface{}, error) {

	// 当前函数只接收struct类型
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr { // 结构体指针
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	out := make(map[string]interface{}, 8)
	queue := make([]interface{}, 0, 1)
	queue = append(queue, in)

	for len(queue) > 0 {
		v := reflect.ValueOf(queue[0])
		if v.Kind() == reflect.Ptr { // 结构体指针
			v = v.Elem()
		}
		queue = queue[1:]
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			vi := v.Field(i)
			//if vi.Kind() == reflect.Ptr { // 内嵌指针
			//	vi = vi.Elem()
			//	if vi.Kind() == reflect.Struct { // 结构体
			//		queue = append(queue, vi.Interface())
			//	} else {
			//		ti := t.Field(i)
			//		if tagValue := ti.Tag.Get("json"); tagValue != "" {
			//			// 存入map
			//			out[tagValue] = vi.Interface()
			//		}
			//	}
			//	break
			//}
			//fmt.Println(vi.Kind())
			if vi.Kind() == reflect.Interface { // 内嵌结构体
				queue = append(queue, vi.Interface())
				break
			}
			// 一般字段
			ti := t.Field(i)
			if tagValue := ti.Tag.Get("json"); tagValue != "" {
				// 存入map
				out[tagValue] = vi.Interface()
			}
		}
	}
	return out, nil
}
