package routes

import (
	"github.com/gin-gonic/gin"
	"boat-service/controllers"
	"boat-service/middleware"
)

func AdminRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	// Protected routes
	adminRoutes := router.Group("/admin").Use(middleware.AuthMiddleware())
	{
		adminRoutes.GET("/users", controllers.ViewUsers)
		adminRoutes.POST("/boats", controllers.AddBoat)
		adminRoutes.PUT("/boats/:id", controllers.EditBoat)
		adminRoutes.DELETE("/boats/:id", controllers.DeleteBoat)
		adminRoutes.GET("/bookings", controllers.ViewBookings)
		adminRoutes.PUT("/bookings/:id", controllers.UpdateBooking)
		adminRoutes.DELETE("/bookings/:id", controllers.CancelBooking)
	}
}
