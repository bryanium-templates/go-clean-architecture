package guitar_core

import (
	"github.com/google/uuid"
)

type Guitar struct {
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
	FindGuitarById(id string) (*Guitar, error)
	UpdateGuitar(id string, guitar Guitar) error
	DeleteGuitar(id string) error
}


