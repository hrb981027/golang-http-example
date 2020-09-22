package config

import (
	"github.com/spf13/viper"
	"log"
)

type Conf struct {
	AppModel string
	AppPort string
}

func LoadConf() Conf {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config文件不存在，开始读取环境变量")

			viper.AutomaticEnv()

			log.Println("环境变量读取成功")
		}
	} else {
		log.Println("配置文件读取成功")
	}

	viper.SetDefault("app_mode", "release")
	viper.SetDefault("app_port", 65535)

	conf := Conf{
		viper.GetString("app_mode"),
		viper.GetString("app_port"),
	}

	return conf
}
