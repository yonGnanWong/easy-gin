package Helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context,err error,data interface{}) {
	code, message := DecodeErr(err)

	//总是返回http状态ok
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
