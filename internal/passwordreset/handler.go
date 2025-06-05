package passwordreset

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RequestPasswordReset (c *gin.Context) {
	log.Println("RequestPasswordReset Endpoint")
}

func VerifyCode (c *gin.Context) {
	log.Println("VerifyCode Endpoint")
}

func UpdatePassword (c *gin.Context) {
	log.Println("UpdatePassword Endpoint")
}
