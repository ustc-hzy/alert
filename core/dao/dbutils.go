package dao

import (
	"alert/core/dao/alert_dao"
	"alert/core/dao/indicator_dao"
	"alert/core/dao/rule_dao"
	"alert/core/dao/task_dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

const (
	userName = "alert_group_lx"
	password = "zkdAlert#"
	ip       = "111.62.122.250"
	port     = "3306"
	dbName   = "data_alert"
)

func InitDB() *gorm.DB {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	DB, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database:%v", err)
	}
	err = DB.AutoMigrate(&indicator_dao.Indicator{}, &rule_dao.Rule{}, &task_dao.Task{}, &alert_dao.Alert{})
	if err != nil {
		log.Fatalf("failed to create the table")
	}
	return DB
}

var DB = InitDB()
