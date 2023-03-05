package initialize

import (
	"MoocShop/internal/dao"
	"fmt"
	"go.uber.org/zap"
)

func Run() {
	// 初始化配置信息
	err := initConfig()
	if err != nil {
		fmt.Printf("init config error: %v\n", err)
	}
	zap.L().Info("init config success")

	// 初始化日志
	err = initLogger()
	if err != nil {
		fmt.Printf("init logger error: %v\n", err)
	}
	zap.L().Info("init logger success")

	// 初始化mysql数据库，
	err = dao.InitMySql()
	if err != nil {
		fmt.Printf("init mysql error: %v\n", err)
	}
	zap.L().Info("init mysql success")

	// 初始化redis
	err = dao.InitRedis()
	if err != nil {
		fmt.Printf("init redis error: ", err)
	}
	zap.L().Info("init redis success")
}
