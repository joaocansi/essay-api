package handler

import "github.com/joaocansi/essay-api/internal/model"

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *CreateUserDTO) ToModel() *model.User {
	return &model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
