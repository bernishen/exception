package repository

import "github.com/bernishen/exception/domain/types/messageScope"

type LoggerRepository interface {
	Log(scope messageScope.MessageScope, code int, msg string)
}
