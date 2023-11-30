package logger

import "github.com/bernishen/exception/domain/types/msgscope"

type LoggerRepository interface {
	Log(scope msgscope.MessageScope, code int, msg string)
}
