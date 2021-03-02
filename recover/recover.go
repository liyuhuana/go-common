package recover

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime/debug"
	"time"

	"github.com/liyuhuana/go-common/utils/file_util"
)

var (
	nEnable = true
	nLogPath string
)

func SetEnable(enable bool) {
	nEnable = enable
}

func SetLogPath(logPath string) {
	nLogPath = logPath
}

func Recover() {
	if !nEnable {
		return
	}

	if e := recover(); e != nil {
		log.Println("Recover => ", e)
		DumpStack()
	}
}

func DumpStack() {
	logPath := path.Join(nLogPath, "recover_log")
	exist, _ := file_util.PathExists(logPath)
	if !exist {
		err := os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	name := path.Join(logPath, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02.15-04-05.000")))
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
		return
	}

	stack := debug.Stack()
	fmt.Fprintln(file, string(stack))
	file.Close()
}
