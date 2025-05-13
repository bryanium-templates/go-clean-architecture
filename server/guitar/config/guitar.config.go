package config

import (
	"server/guitar/controller"
	guitar_core "server/guitar/core"
	"server/guitar/repository"
)

type GuitarFeature struct {
	Controller *controller.GuitarController
}

func InitGuitarFeature() *GuitarFeature {
	repo := repository.NewGuitarRepository()
	service := guitar_core.NewGuitarService(repo)
	controller := controller.NewGuitarController(service)

	return &GuitarFeature{
		Controller: controller,
	}
}