package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"habits/internal/database/models"
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

	err = db.AutoMigrate(models.Habit{}, models.Day{}, models.DayHabit{}, models.HabitWeekDay{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}
