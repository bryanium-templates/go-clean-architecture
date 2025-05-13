package repository

import (
	guitar_core "server/guitar/core"
	"server/infrastructure/database"

	"github.com/google/uuid"
)

type guitarRepository struct{}

func NewGuitarRepository() guitar_core.GuitarRepository {
	return &guitarRepository {}
}

func (r *guitarRepository) AddGuitar(guitar guitar_core.Guitar) error {
	guitar.ID = uuid.New()
	result := database.DB.Create(&guitar)
	return result.Error
}

func (r *guitarRepository) FindAllGuitars() ([]guitar_core.Guitar, error) {
	var guitars []guitar_core.Guitar
	err := database.DB.Find(&guitars).Error
	return guitars, err
}

func (r *guitarRepository) FindGuitarById(id uint) (*guitar_core.Guitar, error) {
	var guitar guitar_core.Guitar
	err := database.DB.First(&guitar, id).Error
	if err != nil {
		return nil, err
	}

	return &guitar, nil
}