package guitar_core

import (
	"errors"
)

type GuitarService struct {
	repository GuitarRepository
}

func NewGuitarService(repo GuitarRepository) *GuitarService {
	return &GuitarService{repository: repo}
}

func (s *GuitarService) CreateGuitar(dto MutateGuitarDTO) error {
	if dto.Category != Acoustic && dto.Category != Electric {
		return errors.New("category must be either 'acoustic' or 'electric'")
	}
	if dto.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	guitar := Guitar{
		Name:        dto.Name,
		Brand:       dto.Brand,
		Description: dto.Description,
		Price:       dto.Price,
		Category:    dto.Category,
	}

	return s.repository.AddGuitar(guitar)
}

func (s *GuitarService) FindAllGuitars() ([]Guitar, error) {
	return s.repository.FindAllGuitars()
}

func (s *GuitarService) FindGuitarById(id string) (*Guitar, error) {
	return s.repository.FindGuitarById(id)
}

func (s *GuitarService) UpdateGuitar(id string, dto MutateGuitarDTO) error {
	if id == "" {
		return errors.New("guitar id is required")
	}
	if dto.Category != Acoustic && dto.Category != Electric {
		return errors.New("category must be either 'acoustic' or 'electric'")
	}
	if dto.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	updated := Guitar{
		Name:        dto.Name,
		Brand:       dto.Brand,
		Description: dto.Description,
		Price:       dto.Price,
		Category:    dto.Category,
	}

	return s.repository.UpdateGuitar(id, updated)
}

func (s *GuitarService) DeleteGuitar(id string) error {
	return s.repository.DeleteGuitar(id)
}
