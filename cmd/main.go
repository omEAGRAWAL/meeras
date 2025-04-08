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
	r.POST("/signup", handlers.SignupHandler)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
