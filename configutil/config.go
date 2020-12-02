package configutil

import (
	"os"
	"path/filepath"

	"git.huoys.com/kit/go-common/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper             //viper结构
	changeHandlerFunc func() //监听到文件修改处理方法
	fileName          string
}

//configType {"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"}
func NewConfig(fileName string,configType string) (*Config, error) {
	conf, err := load(fileName,configType)
	if err != nil {
		return nil, err
	}

	c := &Config{
		Viper:    conf,
		fileName: fileName,
	}

	c.changeMonitor()

	return c, nil
}

func (p *Config) SetChangeHandlerFunc(f func()) {
	p.changeHandlerFunc = f
}

func (p *Config) changeMonitor() {
	//配置文件修改监听
	p.OnConfigChange(func(event fsnotify.Event) {
		if event.Op != fsnotify.Write {
			return
		}
		err1 := p.ReadInConfig()
		if err1 != nil {
			log.CriticalF("config:%v was changed,but read err:", event.Name, err1)
			return
		}
		log.Info("config change:", event.Name)
		if p.changeHandlerFunc != nil {
			p.changeHandlerFunc()
		}
	})
	p.WatchConfig()
}

func load(fileName string, configType string) (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigFile(fileName)
	conf.SetConfigType(configType)
	err := conf.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

//获取app进程运行的绝对路径
func GetAppAbsPath() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("failed get app abs path,err:", err)
		return
	}
	return
}
