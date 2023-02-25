package initialize

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

func initLogger() (err error) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	//
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(viper.GetString("logger.level")))
	if err != nil {
		return
	}
	var core zapcore.Core
	if viper.GetString("app.mode") == "dev" {
		// 如果是dev模式
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}

// getLogWriter 日志切割
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   viper.GetString("logger.filename"), // 日志文件
		MaxSize:    viper.GetInt("logger.max_size"),    // 单个日志最大size
		MaxBackups: viper.GetInt("logger.max_backups"), // 最多备份近多少天的日志，超过则清除旧日志文件
		MaxAge:     viper.GetInt("logger.max_age"),     // 多少天切割一次
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncoder 设置日志展示格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
