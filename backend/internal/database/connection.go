package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"habits/internal/settings"
)

var DB *gorm.DB

func Connect() error {
	envs := settings.GetEnvs().Database
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		envs.Host, envs.User, envs.Database, envs.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}

	DB = db

	return nil
}
