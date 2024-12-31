package router

import (
	"github.com/gin-gonic/gin"
)

func Initilize() {
	// Initialize the router
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)
	
	// Running the server
	router.Run(":3333")
}