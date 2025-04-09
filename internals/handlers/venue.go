package handlers

import (
	"context"
	"meeras/internals/database"
	"meeras/internals/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func VenueHandler(c *gin.Context) {
	var venue models.Venue
	if err := c.ShouldBindJSON(&venue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	venue.ID = primitive.NewObjectID()
	venue.ManagerID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Client.Database("meeras").Collection("venues")

	// Check if user already exists
	res := collection.FindOne(ctx, bson.M{"name": venue.Name})
	if res.Err() == nil { // User already exists
		c.JSON(http.StatusConflict, gin.H{"error": "venue already exists"})
		return
	} else if res.Err() != mongo.ErrNoDocuments { // Other DB errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Insert new user
	_, err1 := collection.InsertOne(ctx, bson.M{
		"_id":         venue.ID,
		"name":        venue.Name,
		"location":    venue.Location,
		"rating":      venue.Rating,
		"description": venue.Description,
		"map_url":     venue.MapURL,
		"manager_id":  venue.ManagerID,
		"packages":    venue.Packages,
	})
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create venue"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Venue registered successfully"})
}

func GetAvailableSlots(c *gin.Context) {
	venueID := c.Param("venueID")
	date := c.Query("date") // YYYY-MM-DD format

	collection := database.Client.Database("meeras").Collection("bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find existing bookings for the venue on the selected date
	cursor, err := collection.Find(ctx, bson.M{"venue_id": venueID, "date": date})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check availability"})
		return
	}

	// Define time slots (Example: 3-hour slots)
	allSlots := []string{"10:00 AM - 1:00 PM", "2:00 PM - 5:00 PM", "6:00 PM - 9:00 PM"}
	bookedSlots := make(map[string]bool)

	for cursor.Next(ctx) {
		var booking models.Booking
		cursor.Decode(&booking)
		bookedSlots[booking.TimeSlot] = true
	}

	// Filter available slots
	availableSlots := []string{}
	for _, slot := range allSlots {
		if !bookedSlots[slot] {
			availableSlots = append(availableSlots, slot)
		}
	}

	c.JSON(http.StatusOK, gin.H{"available_slots": availableSlots})
}

// BookVenue - Allows users to book a venue
func BookVenue(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking details"})
		return
	}

	// Ensure the slot is still available
	collection := database.Client.Database("meeras").Collection("bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existing := collection.FindOne(ctx, bson.M{
		"venue_id":  booking.VenueID,
		"date":      booking.Date,
		"time_slot": booking.TimeSlot,
	})
	if existing.Err() == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Time slot already booked"})
		return
	}

	// Insert booking
	booking.ID = primitive.NewObjectID()
	booking.Status = "Confirmed"
	_, err := collection.InsertOne(ctx, booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book venue"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking successful!"})
}
