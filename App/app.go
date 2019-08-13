package App

import (
	"gin/Service/Database"
	"gin/Routers"
	"gin/Service/Server"

	"github.com/gin-gonic/gin"
)

//创建框架实例,并赋值全局变量R
var R = gin.New()

func Run() {
	//初始化框架配置
	//InitLog()
	initConfig()
	initLog()
	//注册自定义验证
	initValidator()

	Database.InitDb()
	Database.InitRedis()
	Routers.InitRouters(R)
	
	//server start
	s := initServer()
	Server.ListenAndServer(s)

	//注册crontab服务
	initCron()
}
