package exception

import (
	"github.com/bernishen/exception/domain/entity/exception"
	"github.com/bernishen/exception/domain/entity/logger"
	logger2 "github.com/bernishen/exception/domain/repository/logger"
	"github.com/bernishen/exception/domain/service"
	"github.com/bernishen/exception/domain/types/msgscope"
)

func New(scope msgscope.MessageScope, code int, msg string) *exception.Exception {
	ex := exception.NewException(scope, code, msg)
	s := service.ExceptionLoggerService{}
	s.DefaultLoggerRun(ex)
	return ex
}

func Flag(exception *exception.Exception) bool {
	if exception == nil {
		return true
	}
	return false
}

func RegisterLogger(repository logger2.LoggerRepository) {
	if repository == nil {
		return
	}
	logger.DefaultLogger().Register(repository)
}
