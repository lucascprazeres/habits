package habit

import (
	"habits/internal/database/repositories/habit"
	habitsSchema "habits/internal/schemas/habit"
)

type Service interface {
	Create(input habitsSchema.Input) (*habitsSchema.Output, error)
}

type service struct {
	habitsRepo habit.Repository
}

func NewService(repo habit.Repository) Service {
	return &service{
		habitsRepo: repo,
	}
}

func (s service) Create(input habitsSchema.Input) (*habitsSchema.Output, error) {
	return nil, nil
}
