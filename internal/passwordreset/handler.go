package passwordreset

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler Layer Initializer
type Handler struct {
	service *Service
}

func NewHandler (db *gorm.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return &Handler{ service: service}
}

// Handler Layer Methods
func (h *Handler) RequestPasswordReset (c *gin.Context) {
	log.Println("RequestPasswordReset Endpoint")
}

func (h *Handler) VerifyCode (c *gin.Context) {
	log.Println("VerifyCode Endpoint")
}

func (h *Handler) UpdatePassword (c *gin.Context) {
	log.Println("UpdatePassword Endpoint")
}
