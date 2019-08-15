package App

import (
	"gin/Routers"
	"gin/Service/Database"
	"github.com/gin-gonic/gin"
	"os"
)

//创建框架实例,并赋值全局变量R
var R = gin.New()
//日志文件writter
var F *os.File

func Run() {
	//初始化框架配置
	initConfig()
	initLog()
	//注册自定义验证
	initValidator()

	Database.InitDb(F)
	Database.InitRedis()
	Routers.InitRouters(R)

	//server start
	initServer()
	//注册crontab服务
	initCron()
}
