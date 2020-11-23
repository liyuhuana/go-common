package network_http

import (
	"github.com/kataras/iris"
)

type LaunchParam struct {
	Host string
	RelativePath string
	Preprocess iris.Handler
	GetHandlers map[string]iris.Handler
	PostHandlers map[string]iris.Handler
}
