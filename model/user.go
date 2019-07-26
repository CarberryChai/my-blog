package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	PasswordDigest string
	Avatar string `gorm:"size:1000"`
}
