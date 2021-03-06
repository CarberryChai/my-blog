package service

import (
	"golang.org/x/crypto/bcrypt"
	"my-blog/model"
	"my-blog/serializer"
)

// 用户注册
type UserRegister struct {
	UserName        string `form:"userName" json:"userName" binding:"required,min=2,max=30"`
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"passwordConfirm" json:"passwordConfirm" binding:"required,min=8,max=40"`
}

func (user *UserRegister) valid() *serializer.Response {
	if user.Password != user.PasswordConfirm {
		return &serializer.Response{
			Code: 4001,
			Msg:  "输入的两次密码不相等",
		}
	}
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", user.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 4001,
			Msg:  "用户名已注册",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("nickname = ?", user.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 4001,
			Msg:  "昵称已注册",
		}
	}
	return nil
}

func (user *UserRegister) Register() *serializer.Response {
	u := model.User{
		UserName: user.UserName,
		Nickname: user.Nickname,
	}
	if err := user.valid(); err != nil {
		return err
	}
	if err := u.SetPassword(user.Password); err != nil {
		return &serializer.Response{
			Code: 4002,
			Msg:  "加密失败",
		}
	}
	if err := model.DB.Create(&u).Error; err != nil {
		return &serializer.Response{
			Code: 4002,
			Msg:  "注册失败",
		}
	}
	return nil
}

// 用户登录
type UserLogin struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (user *UserLogin) Login() (model.User, *serializer.Response) {
	u := model.User{
		UserName: user.UserName,
	}
	count := 0
	model.DB.Where("user_name = ?", user.UserName).Find(&u).Count(&count)
	if count != 1 {
		return u, &serializer.Response{
			Code: 4004,
			Msg:  "用户不存在",
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(user.Password))
	if err != nil {
		return u, &serializer.Response{
			Code: 4000,
			Msg:  "密码错误",
		}
	}
	return u, nil
}
