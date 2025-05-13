package guitar_core

import "gorm.io/gorm"

type Guitar struct {
	gorm.Model
	Name string
	Brand string
	Price float64
	Category GuitarCategory
}

type GuitarRepository interface {
	AddGuitar(guitar Guitar) error
}


