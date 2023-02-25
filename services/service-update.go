package services

import (
	"crawling/config"
	"crawling/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateById(id string, updatedData utils.DataModel) {
	utils.ValidateData(&updatedData, false)
	update := bson.D{
		{Key: "$set", Value: utils.ToData(updatedData)},
	}
	err := config.DBController.UpdateById(id, update)
	utils.CatchError(err)
}
