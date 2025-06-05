package guitar

import "github.com/gin-gonic/gin"

func RegisterGuitarRoutes(rg *gin.RouterGroup) {
	rg.POST("/add", AddGuitar)
	rg.GET("/", GetAllGuitars)
	rg.PUT("/update/:id", UpdateGuitar)
	rg.DELETE("/delete/:id", DeleteGuitar)
}