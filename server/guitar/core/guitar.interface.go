package guitar_core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Guitar struct {
	gorm.Model
	ID uuid.UUID 
	Name string `gorm:"unique; not null"`
	Brand string
	Description string
	Price float64
	Category GuitarCategory
}

type GuitarRepository interface {
	AddGuitar(guitar Guitar) error
	FindAllGuitars() ([]Guitar, error)
	FindGuitarById(id uint) (*Guitar, error)
	UpdateGuitar(id uint, guitar Guitar) error
	DeleteGuitar(id uint) error
}


