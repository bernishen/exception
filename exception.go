package exception

import (
	"github.com/bernishen/exception/domain/entity"
	"github.com/bernishen/exception/domain/repository"
	"github.com/bernishen/exception/domain/service"
	"github.com/bernishen/exception/domain/types/messageScope"
)

func New(scope messageScope.MessageScope, code int, msg string) *entity.Exception {
	ex := entity.NewException(scope, code, msg)
	s := service.ExceptionLoggerService{}
	s.DefaultLoggerRun(ex)
	return ex
}

func Flag(exception *entity.Exception) bool {
	if exception == nil {
		return true
	}
	return false
}

func RegisterLogger(repository repository.LoggerRepository) {
	if repository == nil {
		return
	}
	entity.DefaultLogger().Register(repository)
}
