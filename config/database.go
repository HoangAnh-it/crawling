package config

import (
	"context"
	"crawling/model"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var (
	DBController *DatabaseController
)

type DatabaseController struct {
	DB *mongo.Database
}

func (db *DatabaseController) ConnectToMongo() {
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

	DBController = &DatabaseController{}
	DBController.DB = client.Database(MONGO_DATABASE)
}

func (db *DatabaseController) InsertOne(data model.Data) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.DB.Collection("malshare").InsertOne(ctx, data)
	if err != nil {
		panic(err)
	}
}

func (db *DatabaseController) InsertMany(data []interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := db.DB.Collection("malshare").InsertMany(ctx, (data))
	if err != nil {
		panic(err)
	}
}
