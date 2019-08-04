package controler

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"my-blog/model"
	"my-blog/serializer"
	"my-blog/service"
	"net/http"
	"strconv"
	"time"
)

type PayLoad struct {
	jwt.Payload
	un string
}

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
	secret := model.Config["secret"]
	now := time.Now()
	hs := jwt.NewHS256([]byte(secret))
	pl := PayLoad{
		Payload: jwt.Payload{
			Issuer:         "carberry",
			ExpirationTime: jwt.NumericDate(now.Add(5 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          strconv.Itoa(int(u.ID)),
		},
		un: u.UserName,
	}

	if token, err := jwt.Sign(pl, hs); err != nil {
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
