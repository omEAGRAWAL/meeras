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

func InsertNewPackageHandler(c *gin.Context) {
	venueName := c.Param("venueName") // Get venue name from URL path
	if venueName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Venue name is required in URL"})
		return
	}

	var newPackage models.Package
	if err := c.ShouldBindJSON(&newPackage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package data"})
		return
	}

	newPackage.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Client.Database("meeras").Collection("venues")

	// Update the venue by pushing the new package
	update := bson.M{
		"$push": bson.M{"packages": newPackage},
	}
	filter := bson.M{"name": venueName}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update venue packages"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Venue not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package added successfully to venue"})
}

func GetAllVenuesHandler(c *gin.Context) {
	// Get the MongoDB collection
	collection := database.Client.Database("meeras").Collection("venues")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query all documents from the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch venues"})
		return
	}
	defer cursor.Close(ctx)

<<<<<<< HEAD
// 	c.JSON(http.StatusOK, gin.H{"message": "Package added successfully to venue"})
// }

//getvenueByname
//create new package
//edit package
//make booking
//update booking
=======
	// Decode all documents into a slice
	var venues []bson.M
	if err := cursor.All(ctx, &venues); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse venues"})
		return
	}

	// Return the data as JSON
	c.JSON(http.StatusOK, gin.H{
		"venues": venues,
	})
}

func UpdatePackageHandler(c *gin.Context) {
	// Extract path parameters
	venueName := c.Param("venueName")
	packageId := c.Param("packageId")

	if venueName == "" || packageId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Venue name and package ID are required in URL"})
		return
	}

	// Convert packageId string to ObjectID
	objID, err := primitive.ObjectIDFromHex(packageId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID format"})
		return
	}

	// Bind new package data from request body
	var updatedPackage models.Package
	if err := c.ShouldBindJSON(&updatedPackage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package data"})
		return
	}

	// Ensure updatedPackage has the correct ID
	updatedPackage.ID = objID

	collection := database.Client.Database("meeras").Collection("venues")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the specific venue and package
	filter := bson.M{
		"name":         venueName,
		"packages._id": objID,
	}

	// Replace the matched package with updated data
	update := bson.M{
		"$set": bson.M{
			"packages.$": updatedPackage,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update package"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Venue or package not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package updated successfully"})
}
>>>>>>> cb83ed7bc4ad0eaff9135dc229944c8696d364e7
