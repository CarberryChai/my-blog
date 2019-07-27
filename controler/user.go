package controler

import (
	"github.com/gin-gonic/gin"
	"my-blog/serializer"
	"my-blog/service"
	"net/http"
)

func Register(ctx *gin.Context) {
	var userInfo service.UserRegister
	if err := ctx.ShouldBind(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userInfo, err := userInfo.Register(); err != nil {
		ctx.JSON(200, err)
	}else {
		ctx.JSON(200, serializer.Response{Data:userInfo, Msg:"注册成功"})
	}
}