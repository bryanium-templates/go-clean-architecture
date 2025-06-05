package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/bryanium-templates/go-clean-architecture/internal/auth"
)

func RegisterRoutes(r *gin.Engine) {
	authHandler := auth.NewHandler(DB)

	authGroup := r.Group("/auth") 
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/signin", authHandler.SignIn)
		authGroup.POST("/signout", authHandler.SignOut)
	}
}