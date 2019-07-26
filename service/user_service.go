package service

import (
	"my-blog/model"
	"my-blog/serializer"
)

type UserRegister struct {
	UserName string `form:"userName" json:"userName" binding:"required,min=8,max=30"`
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"passwordConfirm" json:"passwordConfirm" binding:"required,min=8,max=40"`
}

func (user *UserRegister)Valid()  *serializer.Response{
	if user.Password != user.PasswordConfirm {
		return &serializer.Response{
			Code:4001,
			Msg:"输入的两次密码不想等",
		}
	}
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", user.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:4001,
			Msg:"用户名已注册",
		}
	}
	return nil
}

func (user *UserRegister)Register() (model.User, *serializer.Response){
	u := model.User{
		UserName:user.UserName,
		Nickname:user.Nickname,
	}
	if err := user.Valid(); err != nil {
		return u, err
	}
	if err := u.SetPassword(user.Password); err != nil {
		return u, &serializer.Response{
			Code:4002,
			Msg:"加密失败",
		}
	}
	if err := model.DB.Create(&u).Error; err != nil {
		return u, &serializer.Response{
			Code:4002,
			Msg:"注册失败",
		}
	}
	return u, nil
}
