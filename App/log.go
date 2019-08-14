package App

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)


func initLog()  {
	fName := "./logs/" + time.Now().Format("2006-01-02") + ".log"
	f := pathExist(fName)

	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//设置日志输出文件指针
	log.SetOutput(f)

	//设置gin框架日志处理输入文件指针
	//gin.DefaultWriter = io.MultiWriter(f)

	//将日志在log和控制台中输出
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	F = f

	return
}

//判断目录是否存在,如果不存在则创建并返回log地址
func pathExist(path string) *os.File {
	_ , err := os.Stat("./logs/")
	if os.IsNotExist(err) {
		err = os.Mkdir("./logs/",os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	var f *os.File
	_ , err = os.Stat(path)
	if os.IsNotExist(err) {
		f, _ = os.Create(path)
	}else {
		f ,_ = os.OpenFile(path,os.O_RDWR| os.O_CREATE| os.O_APPEND, 0755)
	}
	return f
}
