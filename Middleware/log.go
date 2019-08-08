package Middleware

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		f := initLog()
		gin.DefaultWriter = io.MultiWriter(f)
		//继续执行接下来的中间件
		c.Next()
	}
}

func initLog() io.Writer{
	file := pathExist("./logs/")
	logfile, _ := os.Create(file)
	return logfile
}

//判断目录是否存在,如果不存在则创建并返回log地址
func pathExist(path string) string {
	_ , err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path,os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return path + time.Now().Format("2006-01-02") + ".log"
}
