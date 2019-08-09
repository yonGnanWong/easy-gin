package App

import (
	"gin/Database"
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
	Database.InitDb()
	Database.InitRedis()
	Routers.InitRouters(R)
	//initCron()

	//server start
	s := InitServer()
	Server.ListenAndServer(s)
}
