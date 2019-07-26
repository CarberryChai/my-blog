package serializer

import "my-blog/model"

type User struct {
	ID uint `json:"id"`
	UserName string `json:"user_name"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	CreatedAt int64 `json:"created_at"`
}
type UserResponse struct {
	Response
	Data User `json:"data"`
}

func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data:User{
			ID:user.ID,
			UserName:user.UserName,
			Nickname:user.Nickname,
			Avatar:user.Avatar,
			CreatedAt:user.CreatedAt.Unix(),
		},
	}
}