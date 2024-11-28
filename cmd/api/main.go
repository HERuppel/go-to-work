package main

import (
	"go-to-work/internal/config"
	"go-to-work/internal/routes"
)

func main() {
	config.Load()

	routes.Initialize()
}
