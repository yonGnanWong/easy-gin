package App

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Server *http.Server

func InitServer() *http.Server {
	//设置gin模式
	gin.SetMode(Config.RunMode)

	Server = &http.Server{
		Addr : Config.Server,
		Handler: R,
	}

	return Server
}
