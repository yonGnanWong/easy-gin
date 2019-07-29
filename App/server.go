package App

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitServer() *http.Server {
	//设置gin模式
	gin.SetMode(Config.RunMode)

	server := &http.Server{
		Addr : Config.Server,
		Handler: R,
	}

	return server
}
