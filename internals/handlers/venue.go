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
