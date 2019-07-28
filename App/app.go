package App

import (
	"gin/Middleware"
	"gin/Routers"
	"gin/Service"
	"github.com/gin-gonic/gin"
)

//创建框架实例,并赋值全局变量R
var R = gin.New()

func Run()  {
	//加载中间件模块
	//终止前端options请求,直接放回
	R.Use(Middleware.Options)
	R.Use(gin.Recovery())

	//初始化框架配置
	InitLog()
	InitConfig()
	InitDb()
	Routers.InitRouters(R)

	//server start
	s := InitServer()
	Service.ListenAndServer(s)
}
