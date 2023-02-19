package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var (
	DB *mongo.Database
)

func ConnectToMongo() {
	USER_NAME := os.Getenv("MONGO_USER_NAME")
	PASSWORD := os.Getenv("MONGO_PASSWORD")
	CLUSTOR := os.Getenv("MONGO_CLUSTOR")
	MONGO_DATABASE := os.Getenv("MONGO_DATABASE")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s.mongodb.net/?retryWrites=true&w=majority", USER_NAME, PASSWORD, CLUSTOR)).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Connecting to your mongodb...")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Failing to connect to mongodb.")
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connect to mongo successfully.")

	DB = client.Database(MONGO_DATABASE)
}
