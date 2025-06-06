package auth

import (
	"log"
	"net/http"

	"github.com/bryanium-templates/go-clean-architecture/utils"
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
		if appErr, ok := err.(*utils.AppError); ok {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(201, gin.H{
		"message":"User signed up successfully", 
		"user": newUser,
	})
}


func (h *Handler) SignIn (c *gin.Context) {
	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error at JSON binding", err)
		c.JSON(400, gin.H{"error": err.Error() })
	}

	token, err := h.service.SignIn(req)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message":"You have signed in successfully",
	})
}

func (h *Handler) SignOut (c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message":"You have logged out successfully",
	})
}