package guitar_core

type GuitarCategory string

const (
	Acoustic GuitarCategory = "acoustic"
	Electric GuitarCategory = "electric"
)

type MutateGuitarDTO struct {
	Name string `json:"name" binding:"required"`
	Brand string `json:"brand" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Category GuitarCategory `json:"category" binding:"required"`
}