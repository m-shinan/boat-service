package main

import (
	"boat-service/database"
	"boat-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.InitDB()

	// Set up router
	router := gin.Default()

	// Define routes
	routes.AdminRoutes(router)
	// routes.UserRoutes(router)
	// routes.BookingRoutes(router)

	// Start server
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
