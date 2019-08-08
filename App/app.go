package App

import (
	"gin/Routers"
	"gin/Service"

	"github.com/gin-gonic/gin"
)

//创建框架实例,并赋值全局变量R
var R = gin.New()

func Run() {
	//初始化框架配置
	//InitLog()
	InitConfig()
	InitDb()
	InitRedis()
	Routers.InitRouters(R)

	//server start
	s := InitServer()
	Service.ListenAndServer(s)
}
