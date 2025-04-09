package main

import (
	"log"
	"meeras/internals/database"
	"meeras/internals/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()

	r := gin.Default()
	// Serve static files (CSS, JS, etc.)
	r.Static("/static", "./static")

	// Serve the HTML file at "/"
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Define routes

	r.POST("/api/signup", handlers.SignupHandler)
	r.POST("/api/login", handlers.LoginHandler)
	r.POST("/api/registervenue", handlers.VenueHandler)
	r.POST("/api/package/:venueName", handlers.InsertNewPackageHandler)
	r.GET("/api/getallvenues", handlers.GetAllVenuesHandler)
	r.GET("/api/updatepackage/:venueName/:packageId", handlers.UpdatePackageHandler)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
