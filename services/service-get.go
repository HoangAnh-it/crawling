package services

import (
	"crawling/config"
	"crawling/model"
	"crawling/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDataById(id string) model.Data {
	objectID, err := primitive.ObjectIDFromHex(id)
	utils.CatchError(err)
	filter := bson.D{{Key: "_id", Value: objectID}}
	data, err := config.DBController.FindOne(filter)
	utils.CatchError(err)
	return *data
}

func GetData(query utils.DataModel) []model.Data {
	result, err := config.DBController.FindAll(query)
	utils.CatchError(err)
	return result
}
