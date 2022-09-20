package log

import (
	"ggl_test/models/dto"
	"ggl_test/utils"
	"github.com/gin-contrib/requestid"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		logger = initLogger()
	}
	return logger
}

func GetLoggerWithCtx(c *dto.AppContext) (l *zap.SugaredLogger) {
	l = GetLogger()
	defer func() {
		if r := recover(); r != nil {
			//	do nothing
		}
	}()
	reqId := requestid.Get(c.GinContext)
	return logger.With("requestId", reqId)
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
