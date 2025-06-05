package guitar

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AddGuitar (c *gin.Context) {
	log.Println("AddGuitar Endpoint")
}

func GetAllGuitars (c *gin.Context) {
	log.Println("GettAllGuitars Endpoint")
}

func UpdateGuitar (c *gin.Context) {
	log.Println("UpdateGuitar Endpoint")
}

func DeleteGuitar (c *gin.Context) {
	log.Println("DeleteGuitar Endpoint")
}