package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(rg *gin.RouterGroup){
	rg.POST("/signup", SignUp)
	rg.POST("/signin", SignIn)
	rg.POST("/signout", SignOut)
}