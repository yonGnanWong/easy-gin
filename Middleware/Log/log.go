package Log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/**
	日志中间件,封装gin框架的日志处理逻辑
 */
func Logger() gin.HandlerFunc {
	//设置gin日志格式
	//return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	return fmt.Sprintf("%s - [%s] \"%s  %s %d %s \"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.StatusCode,
	//		param.ErrorMessage,
	//	)})
}
