package database

import (
	"fmt"

	"github.com/joaocansi/essay-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v database=%v sslmode=disable TimeZone=UTC",
		config.Env.Database.Host,
		config.Env.Database.Port,
		config.Env.Database.User,
		config.Env.Database.Pass,
		config.Env.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
