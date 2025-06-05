package guitar

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler Layer Initializer
type Handler struct {
	service *Service
}

func NewHandler(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return &Handler{ service: service}
}

// Handler Layer Methods
func (h *Handler) AddGuitar (c *gin.Context) {
	log.Println("AddGuitar Endpoint")
}

func (h *Handler) GetAllGuitars (c *gin.Context) {
	log.Println("GetAllGuitars Endpoint")
}

func (h *Handler) UpdateGuitar (c *gin.Context) {
	log.Println("UpdateGuitar Endpoint")
}

func (h *Handler) DeleteGuitar (c *gin.Context) {
	log.Println("DeleteGuitar Endpoint")
}