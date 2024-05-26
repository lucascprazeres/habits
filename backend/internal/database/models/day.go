package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Day struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"primaryKey;uuid"`
	Date      time.Time  `gorm:"unique,not null"`
	DayHabits []DayHabit `gorm:"foreignKey:DayID"`
}

func (Day) TableName() string {
	return "day"
}
