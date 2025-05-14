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

func (r *guitarRepository) UpdateGuitar(id uint, guitar guitar_core.Guitar) error {
	var existing guitar_core.Guitar

	if err := database.DB.First(&existing, id).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&existing).Updates(guitar).Error; err != nil {
		return err
	}

	return nil
}


func (r *guitarRepository) DeleteGuitar(id uint) error {
	if err := database.DB.Delete(&guitar_core.Guitar{}, id).Error; err != nil {
		return err
	}
	return nil
}
