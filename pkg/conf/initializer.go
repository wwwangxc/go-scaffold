package conf

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	confPath string
	initOnce sync.Once
)

// Init ..
func Init() {
	initOnce.Do(func() {
		flag.StringVar(&confPath, "conf", "./config.toml", "default config path")
		flag.Parse()
		initialize()
	})
}

func initialize() {
	confName, confType, confDir := parseConfigPath()
	viper.SetConfigName(confName)
	viper.SetConfigType(confType)
	viper.AddConfigPath(confDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("[%s] 配置文件被修改 -> %s",
			time.Now().Format("2006/01/02 15:04:05"),
			in.Name)
	})
}

func parseConfigPath() (confName, confType, confDir string) {
	lastIndex := strings.LastIndex(confPath, "/")
	if lastIndex == -1 {
		confDir = "./"
	} else {
		confDir = confPath[:lastIndex+1]
	}
	tmp := strings.Split(confPath[lastIndex+1:], ".")
	if len(tmp) < 2 {
		panic(errors.New("wrong config").Error())
	}
	confName = tmp[0]
	confType = tmp[1]
	return
}
