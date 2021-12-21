package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	defaultConfigFilename = "config.yaml"
	configEnvkey          = "JNN_JRS"
)

var Conf Config

type Config struct {
	System systemConf `mapstructure:"system" json:"system" yaml:"system"`
	MySQL  mysqlConf  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type systemConf struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

func init() {
	v := viper.New()
	v.SetConfigFile(getConfigFilePath())
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	unmarshalConfig(v)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) { unmarshalConfig(v) })
}

func unmarshalConfig(v *viper.Viper) {
	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("fail to unmarshal conf file: %s", err))
	}
	initMysql()
}

func getConfigFilePath() (config string) {
	// from command line
	flag.StringVar(&config, "c", "", "input config file path")
	flag.Parse()
	if config != "" {
		fmt.Println("Config file passing from command line:", config)
		return
	}
	// from env var
	if env := os.Getenv(configEnvkey); env != "" {
		config = env
		fmt.Println("Config file passing from environment variable:", config)
		return
	}
	// from default value
	config = defaultConfigFilename
	fmt.Println("Default config file:", config)
	return
}
