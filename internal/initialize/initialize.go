package initialize

import (
	"MoocShop/internal/dao"
	"fmt"
)

func Run() {
	// 初始化配置信息
	err := initConfig()
	if err != nil {
		fmt.Printf("init config error: %v\n", err)
	}

	// 初始化日志
	err = initLogger()
	if err != nil {
		fmt.Printf("init logger error: %v\n", err)
	}

	// 初始化mysql数据库，
	err = dao.InitMySql()
	if err != nil {
		fmt.Printf("init mysql error: %v\n", err)
	}
}
