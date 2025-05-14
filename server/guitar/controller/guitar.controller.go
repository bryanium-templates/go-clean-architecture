package controller

import (
	"net/http"
	guitar_core "server/guitar/core"
	"strconv"

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

func (c *GuitarController) FindAllGuitars(ctx *gin.Context) {
	guitars, err := c.service.FindAllGuitars()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to retrieve guitars"})
		return
	}

	ctx.JSON(http.StatusOK, guitars)
}

func (c *GuitarController) FindGuitarById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Guitar ID"})
		return
	}

	guitar, err := c.service.FindGuitarById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error":"Guitar ID not found"})
		return
	}

	ctx.JSON(http.StatusOK, guitar)
}

func (c *GuitarController) UpdateGuitar(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Guitar ID"})
		return
	}

	var dto guitar_core.MutateGuitarDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.service.UpdateGuitar(uint(id), dto); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Guitar updated successfully"})
}


func (c*GuitarController) DeleteGuitar(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Guitar ID"})
		return
	}

	if err := c.service.DeleteGuitar(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Guitar successfully deleted"})
}