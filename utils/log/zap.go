package log

import (
	"ggl_test/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		logger = initLogger()
	}
	return logger
}

func initLogger() *zap.SugaredLogger {
	var zapLogger *zap.Logger
	if utils.IsProdEnv() {
		zl, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		zapLogger = zl
	} else {
		zl, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		zapLogger = zl
	}
	sugar := zapLogger.Sugar()
	return sugar
}
