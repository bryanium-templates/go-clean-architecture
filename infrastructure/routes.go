package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/bryanium-templates/go-clean-architecture/internal/auth"
	"github.com/bryanium-templates/go-clean-architecture/internal/guitar"
	"github.com/bryanium-templates/go-clean-architecture/internal/passwordreset"
	"github.com/bryanium-templates/go-clean-architecture/internal/user"
)

func RegisterRoutes(r *gin.Engine) {
	auth.RegisterAuthRoutes(r.Group("/auth"))
	user.RegisterUserRoutes(r.Group("/user"))
	passwordreset.RegisterPasswordResetRoutes(r.Group("/passwordreset"))
	guitar.RegisterGuitarRoutes(r.Group("/guitar"))
}