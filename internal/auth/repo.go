package auth

import "gorm.io/gorm"

// Repository Layer Initializer
type Repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *Repository {
	return &Repository{db : db}
}

// Repository Layer Methods