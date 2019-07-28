package Routers

import (
	"gin/Controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(R *gin.Engine)  {
	//404处理
	R.NoRoute(func(c *gin.Context){
		c.String(http.StatusNotFound,"404 not found!")
	})

	//路由设置示例
	R.GET("/", Controllers.GetData)
	R.POST("/post", Controllers.Post)

	return
}
