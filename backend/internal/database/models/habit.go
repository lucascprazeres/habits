package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Habit struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid"`
	Title     string         `gorm:"unique"`
	DayHabits []DayHabit     `gorm:"foreignKey:HabitID"`
	WeekDays  []HabitWeekDay `gorm:"foreignKey:HabitID"`
}

func (Habit) TableName() string {
	return "habit"
}
