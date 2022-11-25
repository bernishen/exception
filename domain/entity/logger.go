package entity

import "github.com/bernishen/exception/domain/repository"

// Logger 日志记录器
type Logger struct {
	Repositories []repository.LoggerRepository
}

var defaultLogger *Logger

// NewLogger 创建一个日至记录器 | Creating a logger
func NewLogger() *Logger {
	logger := &Logger{
		[]repository.LoggerRepository{},
	}
	return logger
}

// DefaultLogger 默认日志记录器 | Get the default logger
func DefaultLogger() *Logger {
	if defaultLogger == nil {
		defaultLogger = NewLogger()
	}
	return defaultLogger
}

// Register 注册日志记录仓储 | Register the loggerRepository
func (l *Logger) Register(loggerRepository repository.LoggerRepository) {
	if loggerRepository == nil {
		return
	}

	l.Repositories = append(l.Repositories, loggerRepository)
}
