package controler

import (
	"github.com/gin-gonic/gin"
	"my-blog/serializer"
	"my-blog/service"
	"net/http"
)

func Register(ctx *gin.Context) {
	var user service.UserRegister
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.Register(); err != nil {
		ctx.JSON(200, err)
	}else {
		ctx.JSON(200, serializer.Response{Msg:"注册成功"})
	}
}

func Login(ctx *gin.Context) {
	var user service.UserLogin
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user, err := user.Login(); err != nil {
		ctx.JSON(200, err)
	}else {
		ctx.JSON(200, serializer.Response{Msg:"登录成功", Data:user})
	}
}