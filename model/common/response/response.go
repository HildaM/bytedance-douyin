package response

/**
 * @Author: Quan
 * @Description:	基本的返回信息封装类
					是所有Response的基类，需要根据具体场景拓展才能使用
 * @File: response
 * @Version: 1.0.0
 * @Date: 2022/5/8 20:50
*/
type BasicResponse struct {
	Code int    `json:"status_code"`
	Msg  string `json:"status_msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

// 创建普通的响应成功
func (BasicResponse) Success() BasicResponse {
	return BasicResponse{
		Code: SUCCESS,
		Msg:  "",
	}
}

// 创建普通响应失败
func (BasicResponse) Fail() BasicResponse {
	return BasicResponse{
		Code: ERROR,
		Msg:  "操作失败",
	}
}

// 创建带信息的响应成功
func (BasicResponse) SuccessWithMsg(msg string) BasicResponse {
	return BasicResponse{
		Code: SUCCESS,
		Msg:  msg,
	}
}

func (BasicResponse) FailWithMsg(msg string) BasicResponse {
	return BasicResponse{
		Code: ERROR,
		Msg:  msg,
	}

}
