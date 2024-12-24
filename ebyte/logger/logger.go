package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level      zapcore.Level
	ErrorFile  string
	AccessFile string
	Encoding   string
}

var logger *zap.Logger

// New 初始化日志记录器
func New(config Config) error {
	zapConfig := zap.Config{
		Encoding:         config.Encoding,
		Level:            zap.NewAtomicLevelAt(config.Level),
		OutputPaths:      []string{"stdout", config.AccessFile}, // 控制台和日志文件路径
		ErrorOutputPaths: []string{"stderr", config.ErrorFile},  // 控制台和错误日志文件路径
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
	}
	var err error
	logger, err = zapConfig.Build()
	if err != nil {
		return err
	}
	return nil
}

func Sync() {
	logger.Sync()
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}
