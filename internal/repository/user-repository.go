package repository

import (
	"fmt"

	"github.com/joaocansi/essay-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if u.db.Where("email = ?", email).First(&user); user.ID == 0 {
		return nil, fmt.Errorf("email not found")
	}
	return &user, nil
}
