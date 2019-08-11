package App

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func initLog()  {
	file := pathExist("./logs/")
	f, _ := os.Create(file)

	//设置gin框架日志处理输入文件指针
	//gin.DefaultWriter = io.MultiWriter(f)

	//将日志在log和控制台中输出
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	return
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
