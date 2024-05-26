package day

import (
	"gorm.io/gorm"
	"habits/internal/database/models"
	"time"
)

type Repository interface {
	GetByDate(date time.Time) (*models.Day, error)
	Create(day *models.Day) (*models.Day, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetByDate(date time.Time) (*models.Day, error) {
	var day models.Day
	err := r.DB.Where("date = ?", date).First(&day).Error
	if err != nil {
		return nil, err
	}

	return &day, nil
}

func (r *repository) Create(day *models.Day) (*models.Day, error) {
	err := r.DB.Create(day).Error
	if err != nil {
		return nil, err
	}
	return day, nil
}
