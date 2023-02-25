package services

import (
	"crawling/config"
	"crawling/utils"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
)

const (
	datePattern = "^[0-9]{4}[-][0-9]{2}[-][0-9]{2}$"
)

var (
	regexpDate = regexp.MustCompile(datePattern)
)

func PushRawString(data utils.RawStringDataModel) string {
	md5 := utils.HashMd5(data.Str)
	sha1 := utils.HashSha1(data.Str)
	sha256 := utils.HashSha256(data.Str)

	var date string
	if data.Date != "" {
		date = data.Date
	} else {
		date = utils.GetCurrentDate()
	}

	dataPosted := utils.DataModel{
		Md5:    md5,
		Sha1:   sha1,
		Sha256: sha256,
		Date:   date,
	}
	existingData, err := config.DBController.FindOne(dataPosted)
	if err != nil && err != mongo.ErrNoDocuments {
		fmt.Println(err)
		panic(err)
	}
	if existingData != nil && existingData.IsSame(utils.ToData(dataPosted)) {
		panic(errors.New("Data is duplicated."))
	}

	id, err := config.DBController.InsertOne(utils.ToData(dataPosted))
	utils.CatchError(err)
	return id
}

func PushData(data utils.DataModel) (id string) {
	utils.ValidateData(&data, true)
	existingData, err := config.DBController.FindOne(data)
	if err != nil && err != mongo.ErrNoDocuments {
		fmt.Println(err)
		panic(err)
	}
	if existingData != nil && existingData.IsSame(utils.ToData(data)) {
		panic(errors.New("Data is duplicated."))
	}
	id, err = config.DBController.InsertOne(utils.ToData(data))
	utils.CatchError(err)
	return id
}
