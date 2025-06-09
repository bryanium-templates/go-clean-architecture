package user

import (
	"log"
	"net/http"

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
	var req UpdateUserRequest 

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": "User not authenticated"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error at JSON binding", err)
		c.JSON(400, gin.H{"error": err.Error() })
	}

	log.Println("Authenticated User", userID)

	c.JSON(200, gin.H {"user":req})
}

func (h *Handler )UpdateProfilePicture(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": "User not authenticated"})
		return
	}

	log.Println("Authenticated User", userID)
}

func (h *Handler )DeleteUser (c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{ "error": "User not authenticated"})
		return
	}

	log.Println("Authenticated User", userID)
}

func (h *Handler) GetAllUsers (c *gin.Context) {
	log.Println("GetAllUsers Endpoint")
}


