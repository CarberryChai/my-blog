package middleware

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"my-blog/serializer"
	"my-blog/util"
	"net/http"
	"os"
	"time"
)
var Pl util.PayLoad
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.PostForm("token")
		if token == "" {
			ctx.JSON(http.StatusOK, serializer.Response{
				Code:4003,
				Msg:"需要登录",
			})
			ctx.Abort()
			return
		}
		var (
			now = time.Now()
			iatValidator = jwt.IssuedAtValidator(now)
			expValidator = jwt.ExpirationTimeValidator(now)
			pl              util.PayLoad
			validatePayload = jwt.ValidatePayload(&pl.Payload, iatValidator, expValidator)
		)
		secret := os.Getenv("SECRET")
		hs := jwt.NewHS256([]byte(secret))
		_, err := jwt.Verify([]byte(token), hs, &pl, validatePayload)
		if err != nil {
			ctx.JSON(http.StatusOK, serializer.Response{
				Code:4003,
				Msg:"token不正确或者已过期，需要登录",
			})
			ctx.Abort()
			return
		}
		Pl = pl
		ctx.Next()
	}
}
