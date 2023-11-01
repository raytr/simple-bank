package repository

import (
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

type BaseRepository interface {
	GetBegin() *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}

func (br *baseRepository) GetBegin() *gorm.DB {
	return br.db.Begin()
}
