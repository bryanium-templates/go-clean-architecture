package passwordreset

import "github.com/gin-gonic/gin"

func RegisterPasswordResetRoutes (rg *gin.RouterGroup) {
	rg.POST("/request",RequestPasswordReset)
	rg.POST("/verify-code", VerifyCode)
	rg.PUT("/update-password", UpdatePassword)
}