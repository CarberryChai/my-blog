package controler

import (
	"github.com/gin-gonic/gin"
	"my-blog/serializer"
	"my-blog/service"
	"my-blog/util"
	"net/http"
)

func Register(ctx *gin.Context) {
	var user service.UserRegister
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.Register(); err != nil {
		ctx.JSON(http.StatusOK, err)
	} else {
		ctx.JSON(http.StatusOK, serializer.Response{Msg: "注册成功"})
	}
}

func Login(ctx *gin.Context) {
	var user service.UserLogin
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := user.Login()
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	if token, err := util.BuildJWT(u); err != nil {
		ctx.JSON(http.StatusOK, serializer.Response{
			Code: 5001,
			Msg:  "登录失败",
		})
	} else {
		ctx.JSON(http.StatusOK, serializer.Response{
			Data: serializer.User{
				UserName: u.UserName,
				Nickname: u.Nickname,
				ID:       u.ID,
				Avatar:   u.Avatar,
				Token:    string(token),
			},
			Msg: "登录成功",
		})
	}
}
