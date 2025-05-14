package guitar_core

import (
	"errors"
)

type GuitarService struct {
	repository GuitarRepository 
}

func NewGuitarService(repo GuitarRepository) *GuitarService {
	return &GuitarService{
		repository: repo,
	}
}

func (s *GuitarService) CreateGuitar(dto MutateGuitarDTO) error {
	if dto.Category != Acoustic && dto.Category != Electric {
		return errors.New("category must be either 'acoustic' or 'electric'")
	}

	if dto.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	guitar := Guitar{
		Name:     dto.Name,
		Brand:    dto.Brand,
		Description: dto.Description,
		Price:    dto.Price,
		Category: dto.Category,
	}

	return s.repository.AddGuitar(guitar)
}

func (s *GuitarService) FindAllGuitars() ([]Guitar, error) {
	return s.repository.FindAllGuitars()
}

func (s *GuitarService) FindGuitarById(id uint) (*Guitar, error) {
	if id == 0 {
		return nil, errors.New("guitar id not found, please provide one")
	}

	return s.repository.FindGuitarById(id)
}

func (s *GuitarService) UpdateGuitar(id uint, dto MutateGuitarDTO) error {
	if id == 0 {
		return errors.New("guitar id not found, please provide one")
	}

	if dto.Category != Acoustic && dto.Category != Electric {
		return errors.New("category must be either 'acoustic' or 'electric'")
	}

	if dto.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	updatedGuitar := Guitar{
		Name:        dto.Name,
		Brand:       dto.Brand,
		Description: dto.Description,
		Price:       dto.Price,
		Category:    dto.Category,
	}

	return s.repository.UpdateGuitar(id, updatedGuitar)
}


func (s *GuitarService) DeleteGuitar(id uint) error {
	return s.repository.DeleteGuitar(id)
}