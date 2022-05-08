package response

import (
	"fmt"
	"testing"
)

func TestTestNormalSuccess(t *testing.T) {
	basicResp := BasicResponse{}.Success()
	msgResp := BasicResponse{}.SuccessWithMsg("成功")
	fmt.Println(basicResp, msgResp)
}
