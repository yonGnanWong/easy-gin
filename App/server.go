package App

import (
	"gin/Service/Server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initServer() {
	//设置gin模式
	gin.SetMode(Config.RunMode)

	server := &http.Server{
		Addr:    Config.Server,
		Handler: R,
	}

	Server.ListenAndServer(server)
	//_ = R.Run(Config.Server)
}
