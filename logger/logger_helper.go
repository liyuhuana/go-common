package logger

import (
	"fmt"
	"github.com/yougg/log4go"
)

var (
	nLogger *Logger
)

func Init(logger *Logger) {
	if logger == nil {
		fmt.Println("Logger init fail, logger is nil!")
		return
	}

	nLogger = logger
}

func Info(args ...interface{}) {
	if nLogger == nil {
		fmt.Println("Logger Info fail, logger is nil, args:", args)
		return
	}
	nLogger.Info(args...)
	logConsoleInfo(log4go.INFO, args...)
}

func Warn(args ...interface{}) {
	if nLogger == nil {
		fmt.Println("Logger Warn fail, logger is nil, args:", args)
		return
	}
	nLogger.Warn(args...)
	logConsoleInfo(log4go.WARNING, args...)
}

func Error(args ...interface{}) {
	if nLogger == nil {
		fmt.Println("Logger Error fail, logger is nil, args:", args)
		return
	}
	nLogger.Error(args...)
	logConsoleInfo(log4go.ERROR, args...)
}

func Critical(args ...interface{}) {
	if nLogger == nil {
		fmt.Println("Logger Critical fail, logger is nil, args:", args)
		return
	}
	nLogger.Critical(args...)
	logConsoleInfo(log4go.CRITICAL, args...)
}
