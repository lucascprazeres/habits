package habit

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"habits/internal/database"
	"habits/internal/database/models"
	"habits/internal/database/repositories/day"
	"habits/internal/database/repositories/dayhabit"
	"habits/internal/database/repositories/habit"
	habitsSchema "habits/internal/schemas/habit"
	"time"
)

type Service interface {
	Create(input habitsSchema.Input) (*habitsSchema.Output, error)
	Toggle(habitID uuid.UUID) (*habitsSchema.Output, error)
}

type service struct {
	habitsRepo    habit.Repository
	daysRepo      day.Repository
	dayHabitsRepo dayhabit.Repository
}

func NewService() Service {
	return &service{
		habitsRepo:    habit.NewRepository(database.DB),
		daysRepo:      day.NewRepository(database.DB),
		dayHabitsRepo: dayhabit.NewRepository(database.DB),
	}
}

func (s service) Create(input habitsSchema.Input) (*habitsSchema.Output, error) {
	habitID := uuid.New()

	var weekDays []models.HabitWeekDay
	for _, weekDay := range input.WeekDays {
		weekDays = append(weekDays, models.HabitWeekDay{
			ID:      uuid.New(),
			HabitID: habitID,
			WeekDay: weekDay,
		})
	}

	habitModel := &models.Habit{
		ID:       habitID,
		Title:    input.Title,
		WeekDays: weekDays,
	}

	err := s.habitsRepo.Create(habitModel)
	if err != nil {
		return nil, err
	}

	output := &habitsSchema.Output{
		ID: habitID.String(),
	}

	return output, nil
}

func (s service) Toggle(habitID uuid.UUID) (*habitsSchema.Output, error) {
	if !s.habitsRepo.ExistsByID(habitID) {
		return nil, errors.New("habit not found")
	}

	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	d, err := s.daysRepo.GetByDate(startOfToday)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		d, err = s.daysRepo.Create(&models.Day{
			ID:   uuid.New(),
			Date: startOfToday,
		})
	}

	if err != nil {
		return nil, err
	}

	dayHabit, err := s.dayHabitsRepo.GetByHabitAndDayIDs(habitID, d.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = s.dayHabitsRepo.Create(&models.DayHabit{
			ID:      uuid.New(),
			HabitID: habitID,
			DayID:   d.ID,
		})

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	if dayHabit != nil {
		if err := s.dayHabitsRepo.Delete(dayHabit.ID); err != nil {
			return nil, err
		}
	}

	return nil, nil
}
