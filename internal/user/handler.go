package user

import (
	"log"

	"github.com/gin-gonic/gin"
)

// User
func UpdateUser (c *gin.Context) {
	log.Println("UpdateUser Endpoint")
}

func UpdateProfilePicture(c *gin.Context) {
	log.Println("UpdateProfilePicture Endpoint")
}

func DeleteUser (c *gin.Context) {
	log.Println("DeleteUser Endpoint")
}

func GetAllUsers (c *gin.Context) {
	log.Println("GetAllUsers Endpoint")
}


// Reset Password Feature
func RequestPasswordReset (c *gin.Context) {
	log.Println("RequestPasswordReset Endpoint")
}

func VerifyCode (c *gin.Context) {
	log.Println("VerifyCode Endpoint")
}

func UpdatePassword (c *gin.Context) {
	log.Println("UpdatePassword Endpoint")
}