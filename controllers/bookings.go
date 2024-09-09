package controllers

import (
	"boat-service/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExample(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Boat Service API",
	})
}

func CreateBooking(c *gin.Context) {
	// Extract data from the request
	var booking struct {
		UserID    int    `json:"user_id"`
		BoatID    int    `json:"boat_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	if err := c.BindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert booking into the database
	_, err := database.DB.Exec(`
		INSERT INTO bookings (user_id, boat_id, start_time, end_time, status)
		VALUES ($1, $2, $3, $4, 'booked')
	`, booking.UserID, booking.BoatID, booking.StartTime, booking.EndTime)

	if err != nil {
		log.Println("Failed to create booking:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully"})
}
