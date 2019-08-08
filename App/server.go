package App

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() *http.Server {
	//设置gin模式
	gin.SetMode(Config.RunMode)

	server := &http.Server{
		Addr:    Config.Server,
		Handler: R,
	}

	return server
}
