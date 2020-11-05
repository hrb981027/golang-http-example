package global

import (
	"github.com/go-redis/redis/v8"
	"golang-http-example/config"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	REDIS  *redis.Client
)
