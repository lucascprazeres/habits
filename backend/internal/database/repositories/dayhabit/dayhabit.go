package dayhabit

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"habits/internal/database/models"
)

type Repository interface {
	GetByHabitAndDayIDs(habitID, dayID uuid.UUID) (*models.DayHabit, error)
	Create(dayHabit *models.DayHabit) error
	Delete(dayHabitID uuid.UUID) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetByHabitAndDayIDs(habitID, dayID uuid.UUID) (*models.DayHabit, error) {
	var dayHabit models.DayHabit
	err := r.DB.Model(&models.DayHabit{}).
		Where("habit_id = ? AND day_id = ?", habitID, dayID).
		Take(&dayHabit).Error
	if err != nil {
		return nil, err
	}
	return &dayHabit, nil
}

func (r *repository) Create(dayHabit *models.DayHabit) error {
	return r.DB.Create(dayHabit).Error
}

func (r *repository) Delete(dayHabitID uuid.UUID) error {
	return r.DB.Delete(&models.DayHabit{}, dayHabitID).Error
}
