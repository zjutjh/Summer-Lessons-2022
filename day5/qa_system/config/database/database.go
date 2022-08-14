package database

import (
	"day5/qa_system/config/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {

	username := config.Config.GetString("database.username")
	password := config.Config.GetString("database.password")
	port := config.Config.GetString("database.port")
	host := config.Config.GetString("database.host")
	name := config.Config.GetString("database.name")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", username, password, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束 提升数据库速度
	})

	if err != nil {
		log.Fatal("DatabaseConnectFailed", err)
	}

	err = autoMigrate(db)
	if err != nil {
		log.Fatal("DatabaseMigrateFailed", err)
	}
	DB = db
}
