package routes

import (
	"github.com/gin-gonic/gin"
	"boat-service/controllers"
	"boat-service/middleware"
)

func UserRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.POST("/forgot-password", controllers.ForgotPassword)

	// Protected routes
	userRoutes := router.Group("/user").Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/profile", controllers.ViewProfile)
		userRoutes.PUT("/profile", controllers.UpdateProfile)
		userRoutes.GET("/boats", controllers.ViewBoatsAndServices)
		userRoutes.POST("/bookings", controllers.BookService)
		userRoutes.GET("/bookings", controllers.ViewBookings)
		userRoutes.DELETE("/bookings/:id", controllers.CancelBooking)
	}
}
