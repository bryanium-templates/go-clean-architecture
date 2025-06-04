package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SignUp (c *gin.Context) {
	log.Println("SignUp Endpoint")
}

func SignIn (c *gin.Context) {
	log.Println("SignIn Endpoint")
}

func SignOut (c *gin.Context) {
	log.Println("SignOut Endpoint")
}