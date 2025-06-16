package logger

import (
	"go.uber.org/zap"
	"sync"
)

var (
	log  *zap.SugaredLogger
	once sync.Once
)

// InitLogger initializes the logger with optional debug mode
func InitLogger(debug bool) {
	once.Do(func() {
		var baseLogger *zap.Logger
		var err error

		if debug {
			baseLogger, err = zap.NewDevelopment()
		} else {
			baseLogger, err = zap.NewProduction()
		}
		if err != nil {
			panic("cannot initialize zap logger: " + err.Error())
		}

		log = baseLogger.Sugar()
	})
}

// L returns the global SugaredLogger instance
func L() *zap.SugaredLogger {
	if log == nil {
		InitLogger(false)
	}
	return log
}

