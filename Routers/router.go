package Routers

import (
	V1 "gin/Api/v1"
	"gin/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(R *gin.Engine) {
	//加载中间件模块
	//终止前端options请求,直接放回
	R.Use(Middleware.Options)
	R.Use(gin.Recovery())
	R.Use(Middleware.Logger())
	//R.Use(Middleware.GlobalMiddleware)

	//404处理
	R.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, http.StatusText(404))
	})

	//路由分组
	v1 := R.Group("api/v1")
	{
		//路由设置示例
		v1.POST("/post", V1.Post)
		v1.GET("/", V1.GetData)
	}

	return
}
