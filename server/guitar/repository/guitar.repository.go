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

func (r *guitarRepository) FindGuitarById(id string) (*guitar_core.Guitar, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var guitar guitar_core.Guitar
	err = database.DB.First(&guitar, "id = ?", parsedID).Error
	if err != nil {
		return nil, err
	}
	return &guitar, nil
}


func (r *guitarRepository) UpdateGuitar(id string, updated guitar_core.Guitar) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var existing guitar_core.Guitar
	if err := database.DB.First(&existing, "id = ?", parsedID).Error; err != nil {
		return err
	}

	return database.DB.Model(&existing).Updates(updated).Error
}



func (r *guitarRepository) DeleteGuitar(id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return database.DB.Delete(&guitar_core.Guitar{}, "id = ?", parsedID).Error
}

