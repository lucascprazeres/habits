package habit

import "gorm.io/gorm"

type Repository interface{}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{DB: db}
}
