package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Load MongoDB URI from environment variable

var Client *mongo.Client

// ConnectDB initializes the MongoDB connection
func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find env file")
		return
	}
	var mongoURI = os.Getenv("mongo_uri")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB Atlas")
}
