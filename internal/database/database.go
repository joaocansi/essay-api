package database

import (
	"fmt"

	"github.com/joaocansi/essay-api/internal/config"
	"github.com/joaocansi/essay-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(DB config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", DB.Host, DB.User, DB.Password, DB.Name, DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return db
}
