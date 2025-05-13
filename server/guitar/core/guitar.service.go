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
		Price:    dto.Price,
		Category: dto.Category,
	}

	return s.repository.AddGuitar(guitar)
}
