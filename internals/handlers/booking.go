// package handlers
//
// import (
//
//	"context"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"meeras/internals/database"
//	"meeras/internals/models"
//	"net/http"
//	"time"
//
// )
//
// func getUserIdByEmail(email string) primitive.ObjectID {
//
//		collection := database.Client.Database("meeras").Collection("users")
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		defer cancel()
//		var user models.User
//		collection.FindOne(ctx, bson.M{
//			"email": email,
//		}).Decode(&user)
//		return user.ID
//	}
//
// func Booking(c *gin.Context) {
//
//	email, ok := c.Get("email")
//	if !ok {
//		fmt.Println("email not found login again")
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": "email not found log in ",
//		})
//	}
//	email1, ok := email.(string)
//
//	UserId := getUserIdByEmail(email1)
//	venueId := c.Param("venuid")
//	packageId := c.Param("packageid")
//
//	OrderId := primitive.ObjectID{}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	collection := database.Client.Database("meeras").Collection("Order")
//
//	type res struct {
//		date         time.Time          `json:"date"`
//		slot         string             `json:"time_slot"`
//		packageid    primitive.ObjectID `json:"package_id"`
//		person_Count int                `json:"person_count"`
//	}
//	var res1 res
//	err := c.ShouldBindJSON(&res1)
//	if err != nil {
//		return
//	}
//	one, err := collection.InsertOne(ctx, bson.M{
//		"_id":          OrderId,
//		"user_id":      UserId,
//		"venue_id":     venueId,
//		"package_id":   packageId,
//		"date":         res1.date,
//		"time_slot":    res1.slot,
//		"status":       "",
//		"person_count": res1.person_Count,
//	})
//	if err != nil {
//		return
//
//	}
//	fmt.Println(one)
//
//	// Check if user already exists
//
//	c.JSON(http.StatusCreated, gin.H{
//		"message": "successsfull",
//	})
//
// }
package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meeras/internals/database"
	"meeras/internals/models"
	"net/http"
	"time"
)

// Error messages as constants
const (
	ErrEmailNotFound     = "email not found, please log in again"
	ErrUserNotFound      = "user not found with the provided email"
	ErrInvalidJSON       = "invalid JSON payload"
	ErrDatabaseOperation = "database operation failed"
	ErrInvalidID         = "invalid ID format"
	MsgBookingSuccessful = "booking created successfully"
)

// BookingRequest represents the JSON payload for booking
type BookingRequest struct {
	Date        time.Time `json:"date" binding:"required"`
	TimeSlot    string    `json:"time_slot" binding:"required"`
	PackageID   string    `json:"package_id,omitempty"`
	PersonCount int       `json:"person_count" binding:"required,min=1"`
}

// getUserIDByEmail retrieves the user ID from the database using email
func getUserIDByEmail(email string) (primitive.ObjectID, error) {
	collection := database.Client.Database("meeras").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&user)

	if err != nil {
		return primitive.NilObjectID, errors.New(ErrUserNotFound)
	}

	return user.ID, nil
}

// Booking handles the creation of a new booking
func Booking(c *gin.Context) {
	// Get email from context (set by authentication middleware)
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": ErrEmailNotFound,
		})
		return
	}

	emailStr, ok := email.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrEmailNotFound,
		})
		return
	}

	// Get user ID from email
	userID, err := getUserIDByEmail(emailStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get venue and package IDs from URL parameters
	venueID := c.Param("venuid")
	packageID := c.Param("packageid")

	// Parse request body
	var bookingReq BookingRequest
	if err := c.ShouldBindJSON(&bookingReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidJSON,
			"details": err.Error(),
		})
		return
	}

	// If packageID is provided in the URL, use it; otherwise use the one from the request
	if packageID == "" && bookingReq.PackageID != "" {
		packageID = bookingReq.PackageID
	}

	// Generate a new ObjectID for the order
	orderID := primitive.NewObjectID()

	// Prepare database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.Client.Database("meeras").Collection("Order")

	// Create booking record
	_, err = collection.InsertOne(ctx, bson.M{
		"_id":          orderID,
		"user_id":      userID,
		"venue_id":     venueID,
		"package_id":   packageID,
		"date":         bookingReq.Date,
		"time_slot":    bookingReq.TimeSlot,
		"status":       "pending", // Default status
		"person_count": bookingReq.PersonCount,
		"created_at":   time.Now(),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseOperation,
			"details": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":  MsgBookingSuccessful,
		"order_id": orderID.Hex(),
	})
}
