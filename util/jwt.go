package util

import (
	"github.com/gbrlsnchs/jwt/v3"
	"my-blog/model"
	"os"
	"strconv"
	"time"
)

type PayLoad struct {
	jwt.Payload
	UN string
}

func BuildJWT(u model.User) (string, error) {
	secret := os.Getenv("SECRET")
	now := time.Now()
	hs := jwt.NewHS256([]byte(secret))
	pl := PayLoad{
		Payload: jwt.Payload{
			Issuer:         "carberry",
			ExpirationTime: jwt.NumericDate(now.Add(1 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          strconv.Itoa(int(u.ID)),
		},
		UN: u.UserName,
	}
	token, err := jwt.Sign(pl, hs)
	if err != nil {
		return "", err
	}
	return string(token), nil
}
