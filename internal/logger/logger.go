package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func Init() {
	logger, _ := zap.NewProduction()
	Logger = logger
}
