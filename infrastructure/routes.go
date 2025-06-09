package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/bryanium-templates/go-clean-architecture/internal/auth"
	"github.com/bryanium-templates/go-clean-architecture/internal/guitar"
	"github.com/bryanium-templates/go-clean-architecture/internal/passwordreset"
	"github.com/bryanium-templates/go-clean-architecture/internal/user"

	"github.com/bryanium-templates/go-clean-architecture/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	authHandler := auth.NewHandler(DB)
	userHandler := user.NewHandler(DB)
	passwordResetHandler := passwordreset.NewHandler(DB)
	guitarHandler := guitar.NewHandler(DB)

	authGroup := r.Group("/auth") 
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/signin", authHandler.SignIn)
		authGroup.POST("/signout", authHandler.SignOut)
	}

	userGroup := r.Group("/user", middleware.JWTAuthMiddleware()) 
	{
		userGroup.PUT("/update", userHandler.UpdateUser)
		userGroup.PUT("/update/profile-picture", userHandler.UpdateProfilePicture)
		userGroup.DELETE("/delete/:id", userHandler.DeleteUser)
		userGroup.GET("/", userHandler.GetAllUsers)
	}

	passwordResetGroup := r.Group("/password-reset")
	{
		passwordResetGroup.POST("/request", passwordResetHandler.RequestPasswordReset)
		passwordResetGroup.POST("/verify", passwordResetHandler.VerifyCode)
		passwordResetGroup.PUT("/update", passwordResetHandler.UpdatePassword)
	}

	guitarGroup := r.Group("/guitar")
	{
		guitarGroup.POST("/add", guitarHandler.AddGuitar)
		guitarGroup.GET("/", guitarHandler.GetAllGuitars)
		guitarGroup.PUT("/update/:id", guitarHandler.UpdateGuitar)
		guitarGroup.DELETE("/delete/:id", guitarHandler.DeleteGuitar)
	}
}