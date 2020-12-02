package configutil

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const (
	describeConfName = "describe.txt"
	personConfName   = "person.txt"
	playerConfName   = "player.txt"
	sourceConfName   = "source.txt"
)

var (
	aaa=`D:\workspace\gopath\src\git.huoys.com\kit\go-common\config`
)


func TestApolloClient_GetConfig(t *testing.T) {
	ac := ApolloClient{
		ClusterName: "default",
		Server:      "http://172.12.12.18:8080",
		AppID:       "coinsafe",
		//ReleaseKey:  "",
		//IP:          "",
		configFiles: sync.Map{},
	}

	config, err := ac.GetConfig("server.toml.txt")
	if err != nil {
		t.Error("failed to get config,err:", err)
	}

	fmt.Println("content:", config)

	config, err = ac.GetConfig("1.json")
	if err != nil {
		t.Error("failed to get config,err:", err)
	}

	fmt.Println("content:", config)
}

func TestApolloClient_GetNotification(t *testing.T) {
	ac := NewApolloClient("http://172.12.12.18:8080", "default", "coinsafe")

	err := ac.AddConfig([]ConfigFile{
		ConfigFile{
			NameSpace:      "server.toml.txt",
			Path:           `D:\workspace\gopath\src\git.huoys.com\kit\go-common\config`,
			notificationID: -1,
		}})
	if err != nil {
		t.Error("failed to add config server.toml.txt,err:", err)
		return
	}

	err = ac.AddConfig([]ConfigFile{ConfigFile{
		NameSpace:      "1.json",
		Path:           `D:\workspace\gopath\src\git.huoys.com\kit\go-common\config`,
		notificationID: -1,
	}})

	if err != nil {
		t.Error("failed to add config 1.json,err:", err)
		return
	}

	err = ac.Start()
	if err != nil {
		t.Error("start apollo,err:", err)
		return
	}

	time.Sleep(time.Second*10)
}
