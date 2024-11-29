package main

import (
	"go-to-work/internal/app"
	"go-to-work/internal/config"
	"go-to-work/internal/routes"
	"log"
)

func main() {
	config.Load()

	container, err := app.NewAppContainer()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v\n", err)
	}

	controllerContainer := &routes.ControllerContainer{
		UserController: container.UserController,
	}

	routes.Initialize(controllerContainer)
}
