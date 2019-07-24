package main

import (
	"log"
	"my-blog/router"
)

func main() {
	app := router.Init()
	log.Fatal(app.Run(":7001"))
}
