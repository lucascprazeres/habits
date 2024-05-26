package models

import "github.com/google/uuid"

type DayHabit struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid"`
	DayID   uuid.UUID `gorm:"type:uuid"`
	HabitID uuid.UUID `gorm:"type:uuid"`
	Day     Day       `gorm:"foreignKey:DayID"`
	Habit   Habit     `gorm:"foreignKey:HabitID"`
}

func (DayHabit) TableName() string {
	return "day_habit"
}
