package logger

import "github.com/yougg/log4go"

func LogInfo(log *Logger, level log4go.Level, args ...interface{}) {
	switch level {
	case log4go.INFO:
		log.Info(args...)
	case log4go.WARNING:
		log.Warn(args...)
	case log4go.ERROR:
		log.Error(args...)
	}
	logConsoleInfo(level, args...)
}
