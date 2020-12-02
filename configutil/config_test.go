package configutil

import (
	"fmt"
	"testing"
	"time"
)

var (
	gameConfig = new(GameConfig)
	testConfig *Config
)

type GameConfig struct {
	GameId  int
	GroupId int
}

func TestNewConfig(t *testing.T) {
	c, err := NewConfig("./test.txt","toml")
	if err != nil {
		t.Error("failed to new config,err:", err)
	}

	testConfig = c

	testConfig.SetChangeHandlerFunc(ChangeHandler)

	fmt.Println("redis addr:",testConfig.Get("redis.Addr"))

	time.Sleep(time.Second*500)
}

func ChangeHandler() {
	fmt.Println("change notify")
	err := testConfig.Viper.UnmarshalKey("service", gameConfig)
	if err != nil {
		fmt.Println("failed to unmarshal key service,err:", err)
		return
	}
	fmt.Println("game config change:", gameConfig)
}
