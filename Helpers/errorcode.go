package Helpers

import "fmt"

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	VALUEERROR        = &Errno{Code: -1, Message: "输入错误"}

	ServerError = &Errno{Code: 10001, Message: "服务器错误"}
	ApiServerError = &Errno{Code: 20001, Message: "接口服务器错误"}
	ModelError = &Errno{Code: 30001, Message: "redis模型错误"}
)

type Errno struct {
	Code int
	Message string
}

//返回错误信息
func (err Errno) Error() string{
	return err.Message
}

//设置 Err 结构体
type Err struct {
	Code int
	Message string
	Err error
}

//声明构造体
func New(errno *Errno,err error) *Err{
	return &Err{Code:errno.Code,Message:errno.Message,Err:err}
}

//添加错误信息
func (err *Err) Add(message string) error{
	err.Message += " " + message
	return err
}

//添加指定格式的错误信息
func (err * Err) Addf(format string,args...interface{}) error{
	err.Message += " " + fmt.Sprintf(format,args...)
	return err
}

//拼接错误信息字符串
func (err *Err) Error() string{
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s",err.Code,err.Message,err.Err)
}

// 解析 错误信息, 返回字符串
func DecodeErr(err error) (int,string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
		case *Err:
			return typed.Code, typed.Message
		case *Errno:
			return typed.Code, typed.Message
		default:
	}
	return ServerError.Code, err.Error()
}
