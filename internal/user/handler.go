package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler Layer Initializer
type Handler struct {
	service *Service
}

func NewHandler (db *gorm.DB) *Handler{
	repo := NewRepository(db)
	service := NewService(repo)
	return &Handler{ service: service}
}


// Handler Layer Methods
func (h *Handler) UpdateUser (c *gin.Context) {
	log.Println("UpdateUser Endpoint")
}

func (h *Handler )UpdateProfilePicture(c *gin.Context) {
	log.Println("UpdateProfilePicture Endpoint")
}

func (h *Handler )DeleteUser (c *gin.Context) {
	log.Println("DeleteUser Endpoint")
}

func (h *Handler) GetAllUsers (c *gin.Context) {
	log.Println("GetAllUsers Endpoint")
}


