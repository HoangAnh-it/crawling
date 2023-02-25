package config

import (
	"context"
	"crawling/model"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DBController *DatabaseController
)

type DatabaseController struct {
	DB        *mongo.Database
	TableName string
}

func (db *DatabaseController) ConnectToMongo() {
	USER_NAME := os.Getenv("MONGO_USER_NAME")
	PASSWORD := os.Getenv("MONGO_PASSWORD")
	CLUSTER := os.Getenv("MONGO_CLUSTER")
	MONGO_DATABASE := os.Getenv("MONGO_DATABASE")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s.mongodb.net/?retryWrites=true&w=majority", USER_NAME, PASSWORD, CLUSTER)).
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
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Connect to mongo successfully.")

	DBController = &DatabaseController{}
	DBController.TableName = "malshare"
	DBController.DB = client.Database(MONGO_DATABASE)
}

func (db *DatabaseController) InsertOne(data model.Data) (id string, err error) {
	result, err := db.DB.Collection(db.TableName).InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (db *DatabaseController) InsertMany(data []interface{}) (ids []string, err error) {
	results, err := db.DB.Collection(db.TableName).InsertMany(context.Background(), data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, id := range results.InsertedIDs {
		ids = append(ids, id.(string))
	}
	return ids, nil
}

func (db *DatabaseController) FindOne(filter interface{}, opts ...*options.FindOneOptions) (data *model.Data, err error) {
	data = &model.Data{}
	err = db.DB.Collection(db.TableName).FindOne(context.Background(), filter, opts...).Decode(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *DatabaseController) FindAll(filter interface{}, opts ...*options.FindOptions) (data []model.Data, err error) {
	ctx := context.Background()
	cur, err := db.DB.Collection(db.TableName).Find(ctx, filter, opts...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cur.Next(ctx) {
		var el model.Data
		_err := cur.Decode(&el)
		if _err != nil {
			fmt.Println(err)
			return nil, err
		}
		data = append(data, el)
	}
	return data, nil
}

func (db *DatabaseController) UpdateById(id string, update interface{}, opts ...*options.UpdateOptions) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	result, err := db.DB.Collection(db.TableName).UpdateByID(context.Background(), objectID, update, opts...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("No document found.")
	}

	return nil
}

func (db *DatabaseController) DeleteById(id string, opts ...*options.DeleteOptions) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	result, err := db.DB.Collection(db.TableName).DeleteOne(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %v", id)
	}
	return nil
}

func (db *DatabaseController) Delete(filter interface{}, opts ...*options.DeleteOptions) error {
	result, err := db.DB.Collection(db.TableName).DeleteMany(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found.")
	}
	return nil
}
