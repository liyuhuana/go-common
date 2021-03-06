package logger

import (
	"github.com/yougg/log4go"
	"time"
)

var (
	consoleLog *log4go.ConsoleLogWriter
)

func init() {
	initConsoleLog()
}

func initConsoleLog() {
	consoleLog = log4go.NewConsoleLogWriter()
	consoleLog.SetFormat("[%T] [%L] %M ")
}

func logConsoleInfo(level log4go.Level, args...interface{}) {
	if consoleLog == nil {
		return
	}

	r := &log4go.LogRecord{
		Level:   level,
		Created: time.Now(),
		Routine: "",
		Source:  "",
		Message: format(args),
	}
	consoleLog.LogWrite(r)
}