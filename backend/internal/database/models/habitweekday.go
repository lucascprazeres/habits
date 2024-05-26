package models

import (
	"github.com/google/uuid"
)

type HabitWeekDay struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid"`
	HabitID uuid.UUID `gorm:"type:uuid"`
	WeekDay uint
	Habit   Habit `gorm:"foreignKey:HabitID"`
}

func (HabitWeekDay) TableName() string {
	return "habit_week_day"
}
