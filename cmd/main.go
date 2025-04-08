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

	log.Println("Server running on :8080")
	r.Run(":8080")
}
