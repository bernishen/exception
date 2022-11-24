package service

import "exception/domain/entity"

type ExceptionLoggerService struct {
}

func (*ExceptionLoggerService) Run(l *entity.Logger, ex *entity.Exception) {
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

func (s *ExceptionLoggerService) DefaultLoggerRun(ex *entity.Exception) {
	s.Run(entity.DefaultLogger(), ex)
}
