package main

import (
	"log"
	"meeras/internals/database"
	"meeras/internals/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()

	r := gin.Default()

	// Define routes
	r.POST("/api/signup", handlers.SignupHandler)
	r.POST("/api/login", handlers.LoginHandler)
	r.POST("/api/registervenue", handlers.VenueHandler)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
