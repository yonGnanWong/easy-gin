package Log

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

/**
	日志中间件,封装gin框架的日志处理逻辑
 */
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//file := pathExist("./logs/")
		//f, _ := os.Create(file)
		// 禁用控制台颜色
		//gin.DisableConsoleColor()

		//gin.DefaultWriter = io.MultiWriter(f)

		// 如果需要将日志同时写入文件和控制台，请使用以下代码
		//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		//继续执行接下来的中间件
		c.Next()
	}
}

//func initLog() io.Writer{
//	//file := pathExist("./logs/")
//	//logfile, _ := os.Create(file)
//	return logfile
//}

//判断目录是否存在,如果不存在则创建并返回log地址
func PathExist(path string) string {
	_ , err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path,os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return path + time.Now().Format("2006-01-02") + ".log"
}
