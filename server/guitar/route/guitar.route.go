package route

import (
	"server/guitar/controller"

	"github.com/gin-gonic/gin"
)

func GuitarRoutes (r *gin.Engine, guitarController *controller.GuitarController) {
	guitars := r.Group("/guitar")
	{
		guitars.POST("", guitarController.AddGuitar)
		guitars.GET("", guitarController.FindAllGuitars)
		guitars.GET("/:id", guitarController.FindGuitarById)
		guitars.PUT("update/:id", guitarController.UpdateGuitar)
		guitars.DELETE("delete/:id",guitarController.DeleteGuitar)
	}
}