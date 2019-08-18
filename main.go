package main

import (
	"github.com/joho/godotenv"
	"log"
	"my-blog/cache"
	"my-blog/model"
	"my-blog/router"
)

func main() {
	// 加载配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// 连接mysql
	model.SetupDatabase()
	// 连接redis
	cache.Init()
	// 装载路由
	app := router.Init()
	// 启动app
	log.Fatal(app.Run(":7001"))
}
