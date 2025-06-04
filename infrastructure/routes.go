package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/bryanium-templates/go-clean-architecture/internal/auth"
)

func RegisterRoutes(r *gin.Engine) {
	auth.RegisterAuthRoutes(r.Group("/auth"))
}