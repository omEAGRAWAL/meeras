package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"meeras/internals/database"
	"meeras/internals/handlers"
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()

	r := gin.Default()
	// Serve static files (CSS, JS, etc.)
	r.Static("/static", "./static")

	// Serve the HTML file at "/"
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})
	// Define routes

	r.POST("/api/signup", handlers.SignupHandler)
	r.POST("/api/login", handlers.LoginHandler)
	r.POST("/api/registervenue", handlers.VenueHandler)
	r.POST("/api/package/:venueName", handlers.InsertNewPackageHandler)
	r.GET("/api/getallvenues", handlers.GetAllVenuesHandler)

	log.Println("Server running on :8080")
	err := r.Run(":8080")
	if err != nil {
		return 
	}
}
