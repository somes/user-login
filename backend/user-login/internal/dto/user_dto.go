package dto

import "server/internal/model"

type UserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}
