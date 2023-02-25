package main

import (
	"MoocShop/internal/initialize"
	"MoocShop/internal/router"
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	// 初始化配置文件信息
	initialize.Run()

	// 初始化路由信息
	engine := router.SetupRouter()
	engine.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
}
