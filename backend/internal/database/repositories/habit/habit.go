package habit

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"habits/internal/database/models"
)

type Repository interface {
	Create(habit *models.Habit) error
	ExistsByID(habitID uuid.UUID) bool
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{DB: db}
}

func (r repository) ExistsByID(habitID uuid.UUID) bool {
	err := r.DB.First(&models.Habit{}, "id = ?", habitID).Error
	return err == nil
}

func (r repository) Create(habit *models.Habit) error {
	return r.DB.Create(habit).Error
}
