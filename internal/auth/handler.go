package auth

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
func (h *Handler) SignUp (c *gin.Context) {

}


func (h *Handler) SignIn (c *gin.Context) {
	log.Println("SignIn Endpoint")
}

func (h *Handler) SignOut (c *gin.Context) {
	log.Println("SignOut Endpoint")
}