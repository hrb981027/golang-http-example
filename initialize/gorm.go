package initialize

import (
	"fmt"
	"golang-http-example/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() {
	switch global.CONFIG.System.DbType {
	case "mysql":
		GormMysql()
	default:
		GormMysql()
	}
}

func GormMysql() {
	var err error

	conf := global.CONFIG.Mysql

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		conf.Username,
		conf.Password,
		conf.HostName,
		conf.HostPort,
		conf.Database,
	)

	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败")
	}
}
