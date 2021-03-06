package network_http

import (
	"github.com/kataras/iris"

	"github.com/liyuhuana/go-common/logger"
)

func StartHttpServer(param LaunchParam) {
	app := iris.Default()

	// 注册监听
	party := app.Party(param.RelativePath, param.Preprocess)
	for k, v := range param.GetHandlers {
		party.Get(k, v)
	}

	for k, v := range param.PostHandlers {
		party.Post(k, v)
	}

	// 启动
	hostAddr := param.Host
	err := app.Run(iris.Addr(hostAddr))
	if err != nil {
		logger.Info("Server start fail, host [", hostAddr, "], error:", err)
	} else {
		logger.Error("Server start >>> host [", hostAddr, "]")
	}
}
