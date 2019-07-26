package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string
	Nickname string
	PasswordDigest string
	Avatar string `gorm:"size:1000"`
}
const PassWordCost = 12
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return  err
	}
	user.PasswordDigest = string(bytes)
	return nil
}