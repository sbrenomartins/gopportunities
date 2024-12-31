package main

import (
	"github.com/sbrenomartins/gopportunities/config"
	"github.com/sbrenomartins/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Initialize()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Initialize the router
	router.Initilize()
}
