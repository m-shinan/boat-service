package routes

import (
	"boat-service/controllers"
	"github.com/gin-gonic/gin"
)

func BookingsRoutes(router *gin.Engine) {
	router.GET("/admin/example", controllers.GetExample)
	// Add other admin routes here
}