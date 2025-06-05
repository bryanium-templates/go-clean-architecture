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
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error at JSON binding", err)
		c.JSON(400, gin.H{"error": err.Error() })
	}

	newUser, token, err := h.service.SignUp(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to sign up user"})
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(201, gin.H{
		"message":"User signed up successfully", 
		"user": newUser,
	})
}


func (h *Handler) SignIn (c *gin.Context) {
	log.Println("SignIn Endpoint")
}

func (h *Handler) SignOut (c *gin.Context) {
	log.Println("SignOut Endpoint")
}