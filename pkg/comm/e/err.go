//@Author: wulinlin
//@Description:
//@File:  err
//@Version: 1.0.0
//@Date: 2023/12/01 22:57

package e

var MessageFlag = map[int]string{
	SUCCESS:          "ok",
	ERROR:            "fail",
	BAD_INPUT_PARAMS: "请求参数错误",
	NOT_LOGIN:        "用户未登录",
}

func GetMessage(code int) string {
	msg, ok := MessageFlag[code]
	if ok {
		return msg
	}
	return MessageFlag[ERROR]
}
