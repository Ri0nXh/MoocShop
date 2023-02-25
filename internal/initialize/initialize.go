package initialize

import "fmt"

func Run() {
	// 初始化配置信息
	err := initConfig()
	if err != nil {
		fmt.Printf("init config error: %v\n", err)
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("init logger error: %v\n", err)
	}

}
