package user

import (
	"log"

	"github.com/gin-gonic/gin"
)

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


