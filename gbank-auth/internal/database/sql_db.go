package database

import (
	"github.com/guilhermealvess/gbank/gbank-auth/internal/repository"
	"gorm.io/gorm"
)

type SqlDB[T DatabaseEntities] struct {
	repository.SqlDB
	db *gorm.DB
}

func NewSqlDB[T DatabaseEntities](db *gorm.DB) *SqlDB[T] {
	return &SqlDB[T]{
		db: db,
	}
}

func (s *SqlDB[T]) Insert(model T) error {
	// Migrate the schema
	// s.db.AutoMigrate(&Product{})

	s.db.Create(model)

	return nil
}
