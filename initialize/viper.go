package initialize

import (
	"github.com/spf13/viper"
	"golang-http-example/global"
	"log"
)

func Viper() {
	v := viper.New()

	v.AddConfigPath("./")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config文件不存在，开始读取环境变量")

			v.AutomaticEnv()

			log.Println("环境变量读取成功")
		}
	} else {
		log.Println("配置文件读取成功")
	}

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic("配置读取失败")
	}
}
