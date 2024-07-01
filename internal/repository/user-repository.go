package repository

import (
	"github.com/joaocansi/essay-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user *model.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
