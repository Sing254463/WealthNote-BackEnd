package database

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB() {
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("MONGODB_URI environment variable is not set")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)

    var err error
    MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = MongoClient.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    fmt.Println("Connected to MongoDB!")
}

func DisconnectMongoDB() {
    if err := MongoClient.Disconnect(context.TODO()); err != nil {
        log.Fatalf("Failed to disconnect from MongoDB: %v", err)
    }
    fmt.Println("Disconnected from MongoDB!")
}