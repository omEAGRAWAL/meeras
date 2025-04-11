package main

import (
	"log"
	"meeras/internals/config"
	"meeras/internals/database"
	"meeras/internals/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()
	config.InitCloudinary()
	r := gin.Default()
	// Serve static files (CSS, JS, etc.)
	// r.Static("/static", "./static")
	r.Static("/assets", "./internals/client/dist/assets")

	// Catch-all route to serve index.html (for SPA routes like /login, /dashboard, etc.)
	r.NoRoute(func(c *gin.Context) {
		c.File("./internals/client/dist/index.html")
	})

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
	r.DELETE("/api/deletepackage/:packageId", handlers.DeletePackageHandler)
	r.GET("/api/getallpackages", handlers.GetAllPacakages)
	r.POST("/api/updatevenue/:venueId", handlers.UpdateVenue)
	r.DELETE("/api/deletevenue/:venueId", handlers.DeleteVenue)

	log.Println("Server running on :8089")
	r.Run(":8089")
}
