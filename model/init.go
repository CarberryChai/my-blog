package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB

func SetupDatabase() {
	mysqlConf := os.Getenv("MYSQL_CONF")
	db, err := gorm.Open("mysql", mysqlConf)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(50)
	DB = db
	migration()
}
