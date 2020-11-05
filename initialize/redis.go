package initialize

import (
	"github.com/go-redis/redis/v8"
	"golang-http-example/global"
)

func Redis() {
	conf := global.CONFIG.Redis

	global.REDIS = redis.NewClient(&redis.Options{
		Addr:     conf.HostName + ":" + conf.HostPort,
		Password: conf.Password,
		DB:       0,
	})
}
