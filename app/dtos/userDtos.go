package dtos

import "github.com/ZaphCode/F-SR-ChatApp/domain"

type NewUserDto struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func (d NewUserDto) AdaptToUser() domain.User {
	return domain.User{
		Username: d.Username,
		Email:    d.Email,
		Password: d.Password,
	}
}
