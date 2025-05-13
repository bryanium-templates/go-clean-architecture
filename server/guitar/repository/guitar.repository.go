package repository

import (
	guitar_core "server/guitar/core"
	"server/infrastructure/database"
)

type guitarRepository struct{}

func NewGuitarRepository() guitar_core.GuitarRepository {
	return &guitarRepository {}
}

func (r *guitarRepository) AddGuitar(guitar guitar_core.Guitar) error {
	result := database.DB.Create(&guitar)
	return result.Error
}
