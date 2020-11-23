package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var SugarLogger *zap.SugaredLogger // SugarLogger需要全局使用记录日志

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	// zapcore.DebugLevel: 日志级别
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller()) // zap.AddCaller(): 记录调用函数
	SugarLogger = logger.Sugar()             // 通过主logger的方法获取SugarLogger
}

// getEncoder获取编码器(如何写入日志)
// 时间使用人类可读的方式
// 使用大写字母标识日志级别
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 指定日志输出
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("logfile/zap.log")
	return zapcore.AddSync(file)
}
