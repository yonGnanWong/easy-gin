package Helpers

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	VALUEERROR        = &Errno{Code: -1, Message: "输入错误"}
	InternalServerError = &Errno{Code: 10001, Message: "服务器错误"}
	ApiServerError = &Errno{Code: 20001, Message: "接口服务器错误"}
)

type Errno struct {
	Code int
	Message string
}
