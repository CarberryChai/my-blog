package router

import (
	"github.com/gin-gonic/gin"
	"my-blog/controler"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})
	// login
	router.POST("/login", controler.UserLogin)
	return router
}
