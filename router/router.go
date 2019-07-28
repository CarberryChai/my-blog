package router

import (
	"github.com/gin-gonic/gin"
	"my-blog/controler"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello World!")
		})
		// login
		api.POST("/register", controler.Register)
		api.POST("/login", controler.Login)
	}
	return router
}
