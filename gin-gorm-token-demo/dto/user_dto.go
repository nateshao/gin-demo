package dto

import "gin-gorm-token-demo/model"

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToUserDto(user *model.User) UserDto {
	return UserDto{
		Username: user.Username,
		Password: user.Password,
	}
}
