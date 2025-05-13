package controller

import (
	"net/http"
	guitar_core "server/guitar/core"

	"github.com/gin-gonic/gin"
)

type GuitarController struct {
	service *guitar_core.GuitarService 
}

func NewGuitarController(service *guitar_core.GuitarService) *GuitarController {
	return &GuitarController { service }
}

func (c *GuitarController) AddGuitar(ctx *gin.Context) {
	var guitarData guitar_core.MutateGuitarDTO
	if err := ctx.ShouldBindJSON(&guitarData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateGuitar(guitarData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message":"A guitar has been added successfully", "data": guitarData})
}