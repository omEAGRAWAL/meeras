package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"meeras/cloudinary"
	"meeras/internals/database"
	"meeras/internals/handlers"
	"meeras/internals/middleware"
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()

	r := gin.Default()
	// Serve static files (CSS, JS, etc.)
	r.Static("/static", "./static")

	// Serve the HTML file at "/"
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Define routes

	r.POST("/api/signup", handlers.SignupHandler)
	r.POST("/api/login", handlers.LoginHandler)
	r.POST("/api/registervenue", middleware.Authmiddle(), handlers.VenueHandler)
	r.POST("/api/package/:venueName", handlers.InsertNewPackageHandler)
	r.GET("/api/getallvenues", handlers.GetAllVenuesHandler)
	r.GET("/api/updatepackage/:venueName/:packageId", handlers.UpdatePackageHandler)
	r.POST("/api/upload/image", middleware.Authmiddle(), cloudinary.UploadHandler)
	r.POST("/api/booking/:venuid/:packageid", middleware.Authmiddle(), handlers.Booking)

	log.Println("Server running on :8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
