package controler

import (
	"github.com/gin-gonic/gin"
	"my-blog/service"
	"net/http"
)

func Login(ctx *gin.Context) {
	var user service.Login
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data":user})
}