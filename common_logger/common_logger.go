package common_logger

import "fmt"

var (
	nLogInfo, nLogWarn, nLogError func(...interface{})
)

func Init(logInfo, logWarn, logError func(...interface{})) {
	nLogInfo = logInfo
	nLogWarn = logWarn
	nLogError = logError
}

func LogInfo(args...interface{}) {
	if nLogInfo != nil {
		nLogInfo(args...)
	} else {
		fmt.Println(args...)
	}
}

func LogWarn(args...interface{}) {
	if nLogWarn != nil {
		nLogWarn(args...)
	} else {
		fmt.Println(args...)
	}
}

func LogError(args...interface{}) {
	if nLogError != nil {
		nLogError(args...)
	} else {
		fmt.Println(args...)
	}
}
