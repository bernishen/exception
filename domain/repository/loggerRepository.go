package repository

import (
	"exception/domain/types/messageScope"
)

type LoggerRepository interface {
	Log(scope messageScope.MessageScope, code int, msg string)
}
