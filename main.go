package main

import (
	"MoocShop/internal/router"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

const (
	CONFIG_PATH = "./configs"
)

func main() {

	// 初始化配置文件信息
	InitConfig()

	// 初始化路由信息
	engine := router.SetupRouter()
	engine.Run(fmt.Sprintf("0.0.0.0:%d", viper.GetInt("port")))
}

func InitConfig() {
	files, err := ioutil.ReadDir(CONFIG_PATH)
	if err != nil {
		fmt.Println("config dir load error: ", err)
	}
	viper.SetConfigType("yaml")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := os.Open(fmt.Sprintf("%s/%s", CONFIG_PATH, file.Name()))
		if err != nil {
			fmt.Printf("open %s error", file)
		}
		err = viper.MergeConfig(f)
		if err != nil {
			fmt.Println("viper merge config error: ", err)
		}
	}
}
