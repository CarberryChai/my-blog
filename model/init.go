package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
)

type JSON map[string]string

var DB *gorm.DB
var Config JSON

func SetupDatabase() {
	content, err := ioutil.ReadFile("env.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var data JSON
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err.Error())
	}
	Config = data
	db, err := gorm.Open("mysql", data["mysql"])
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(50)
	DB = db
	migration()
}
