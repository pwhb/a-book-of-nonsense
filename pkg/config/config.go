package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeClient() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, clientOptions)
}
