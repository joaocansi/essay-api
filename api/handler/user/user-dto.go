package handler

import "github.com/joaocansi/essay-api/internal/model"

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=24"`
}

func (c *CreateUserDTO) ToModel() *model.User {
	return &model.User{
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
	}
}
