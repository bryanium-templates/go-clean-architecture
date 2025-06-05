package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/bryanium-templates/go-clean-architecture/internal/auth"
	"github.com/bryanium-templates/go-clean-architecture/internal/guitar"
)

func RegisterRoutes(r *gin.Engine) {
	authHandler := auth.NewHandler(DB)
	guitarHandler := guitar.NewHandler(DB)

	authGroup := r.Group("/auth") 
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/signin", authHandler.SignIn)
		authGroup.POST("/signout", authHandler.SignOut)
	}

	guitarGroup := r.Group("/guitar")
	{
		guitarGroup.POST("/add", guitarHandler.AddGuitar)
		guitarGroup.GET("/", guitarHandler.GetAllGuitars)
		guitarGroup.PUT("/update/:id", guitarHandler.UpdateGuitar)
		guitarGroup.DELETE("/delete/:id", guitarHandler.DeleteGuitar)
	}
}