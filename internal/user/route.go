package user

import "github.com/gin-gonic/gin"


func RegisterUserRoutes (rg *gin.RouterGroup) {
	rg.PUT("/update", UpdateUser)
	rg.PUT("/update/profile-picture", UpdateProfilePicture)
	rg.DELETE("/delete", DeleteUser)

	rg.POST("/request-password-reset",RequestPasswordReset)
	rg.POST("/verify-code", VerifyCode)
	rg.PUT("/update/password", UpdatePassword)
}