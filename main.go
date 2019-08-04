package main

import (
	"log"
	"my-blog/cache"
	"my-blog/model"
	"my-blog/router"
)

func main() {
	// 连接mysql
	model.SetupDatabase()
	// 连接redis
	cache.Init()
	// 装载路由
	app := router.Init()
	// 启动app
	log.Fatal(app.Run(":7001"))
}
