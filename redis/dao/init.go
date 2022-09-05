package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DSN string = "demo:demo@tcp(192.168.1.102:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	})
	if err != nil {
		log.Fatal(err)
	}
}
