package httpserv

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/myste1tainn/msnet"
	"github.com/spf13/viper"
)

func Run() {
	a := gin.Default()
	a.Use(msnet.TraceWithDefaultOptions())

	bindCreateGameRoute(a)
	bindHealthRoute(a)
	appPort := viper.GetString("app.port")
	if appPort == "" {
		appPort = "1323"
	}
	a.Run(":" + appPort)
	//server := a.Run(":" + appPort)
	//a.RegisterShutdownListener(server, a.RegisterListeners(), exCloseFunc)
}

func exCloseFunc() {
	fmt.Println("Close function run here")
}
