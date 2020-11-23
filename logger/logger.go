package logger

import (
	"github.com/liyuhuana/go-common/utils/file_util"
	"fmt"
	"github.com/yougg/log4go"
	"os"
	"path/filepath"
	"strings"
)

type Logger struct {
	log log4go.Logger
}

func NewLogger(logPath, logName string) (*Logger, error) {
	exist, err := file_util.PathExists(logPath)
	if err != nil {
		return nil, err
	}

	// 如果文件夹不存在，则要新创建文件夹
	if !exist {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	logName = filepath.Join(logPath, logName)

	logWriter := log4go.NewFileLogWriter(logName, true)
	logWriter.SetFormat("[%T] [%L] (%I) (%s) %M ")
	logWriter.SetRotateSize(50 * 1024 * 1024)
	logWriter.SetRotateDaily(true)
	logWriter.SetMaxDays(10)

	log4go.EnableGoRoutineID = true

	v := log4go.Logger{}
	v.AddFilter("file", log4go.DEBUG, logWriter)
	return &Logger{log:v}, nil
}

func (p *Logger) Debug(args ...interface{}) {
	p.log.Debug(format(args...))
}

func (p *Logger) Info(args ...interface{}) {
	p.log.Info(format(args...))
}

func (p *Logger) Warn(args ...interface{}) {
	p.log.Warn(format(args...))
}

func (p *Logger) Error(args ...interface{}) {
	p.log.Error(format(args...))
}

func (p *Logger) Critical(args ...interface{}) {
	p.log.Critical(format(args...))
}

func (p *Logger) DebugF(format string, args ...interface{}) {
	p.log.Debug(format, args...)
}

func (p *Logger) InfoF(format string, args ...interface{}) {
	p.log.Info(format, args...)
}

func (p *Logger) WarnF(format string, args ...interface{}) {
	p.log.Warn(format, args...)
}

func (p *Logger) ErrorF(format string, args ...interface{}) {
	p.log.Error(format, args...)
}

func (p *Logger) CriticalF(format string, args ...interface{}) {
	p.log.Critical(format, args...)
}

func format(args ...interface{}) (msg string) {
	msg = fmt.Sprintf(strings.Repeat("%v", len(args)), args...)
	return msg
}