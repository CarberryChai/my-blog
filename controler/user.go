package controler

import (
	"github.com/gin-gonic/gin"
	"my-blog/model"
	"net/http"
)

func UserLogin(ctx *gin.Context) {
	var user model.Login
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data":user})
}