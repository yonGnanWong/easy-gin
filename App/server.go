package App

import (
	"gin/Service/Server"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func initServer() {
	//设置gin模式
	gin.SetMode(Config.RunMode)

	server := &http.Server{
		Addr:    Config.Server,
		Handler: R,
	}
	sysType := runtime.GOOS

	if sysType == "linux" {
		Server.ListenAndServer(server)
	}else if sysType == "windows" {
		_ = R.Run(Config.Server)
	}
}
