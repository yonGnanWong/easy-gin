package Routers

import (
	"gin/Api/v1"
	"gin/Middleware/Header"
	"gin/Middleware/Log"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitRouters(R *gin.Engine) {
	//加载中间件模块
	//终止前端options请求,直接放回
	R.Use(Header.Options)
	//R.Use(Middleware.Logger())
	R.Use(gin.Recovery())


	R.Use(gin.Logger())
	//R.Use(Middleware.GlobalMiddleware)

	//404处理
	R.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, http.StatusText(404))
	})

	//路由分组
	r1 := R.Group("api/v1")
	{
		//路由设置示例
		r1.POST("/post", v1.C)
		r1.GET("/", v1.R)
	}

	return
}
