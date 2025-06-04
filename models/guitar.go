package models

import (
	"time"

	"github.com/google/uuid"
)
type GuitarCategory string

type Guitar struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid"`

	Name string `gorm:"type:varchar(100);not null"`
	Brand string `gorm:"type:varchar(100);not null"`

	Description string `gorm:"type:text"`
	Price float64 `gorm:"type:decimal(10,2);not null"`

	Category GuitarCategory `gorm:"type:varchar(50);not null;column:category"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}


const (
	Acoustic GuitarCategory = "acoustic"
	Electric GuitarCategory = "electric"
	Bass GuitarCategory = "bass"
)