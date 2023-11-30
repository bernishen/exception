package service

import (
	"github.com/bernishen/exception/domain/entity/exception"
	"github.com/bernishen/exception/domain/entity/logger"
)

type ExceptionLoggerService struct {
}

func (*ExceptionLoggerService) Run(l *logger.Logger, ex *exception.Exception) {
	if l == nil || l.Repositories == nil || len(l.Repositories) == 0 {
		return
	}
	if ex == nil {
		return
	}

	for _, logger := range l.Repositories {
		if logger == nil {
			continue
		}
		logger.Log(ex.Scope, ex.Code, ex.Message)
	}
}

func (s *ExceptionLoggerService) DefaultLoggerRun(ex *exception.Exception) {
	s.Run(logger.DefaultLogger(), ex)
}
