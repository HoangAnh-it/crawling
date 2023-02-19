package services

import (
	"context"
	"crawling/config"
	"crawling/model"
	"crawling/utils"
	"errors"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	datePattern = "^[0-9]{4}[-][0-9]{2}[-][0-9]{2}$"
)

var (
	regexpDate = regexp.MustCompile(datePattern)
)

/*
 Insert data which is located in folder "output"
*/
func InsertData(date string, typeFile string) (err error) {
	defer func() {
		if _err := recover(); _err != nil {
			err = _err.(error)
		}
	}()

	if !regexpDate.MatchString(date) {
		return errors.New(`Incompatible date format. Must be "yyyy-MM-dd".\n`)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch typeFile {
	case "md5", "sha1", "sha256", "base64":
		filePath := filepath.Join("output", strings.ReplaceAll(date, "-", "/"), typeFile+".txt")
		content, err := utils.ReadFile(filePath)
		utils.CatchError(err)
		_, err = config.DB.Collection("malshare").InsertOne(ctx, model.Data{
			Type:     typeFile,
			HashCode: content,
			Date:     date,
		})
		utils.CatchError(err)
		break

	case "all", "":
		subFilePath := filepath.Join("output", strings.ReplaceAll(date, "-", "/"))
		md5, errMd5 := utils.ReadFile(subFilePath + "/md5.txt")
		sha1, errSha1 := utils.ReadFile(subFilePath + "/sha1.txt")
		sha256, errSha256 := utils.ReadFile(subFilePath + "/sha256.txt")
		base64, errBase64 := utils.ReadFile(subFilePath + "/base64.txt")
		utils.CatchError(errMd5)
		utils.CatchError(errSha1)
		utils.CatchError(errSha256)
		utils.CatchError(errBase64)
		config.DB.Collection("malshare").InsertMany(ctx,
			[]interface{}{
				*model.CreateDate("md5", md5, date),
				*model.CreateDate("sha1", sha1, date),
				*model.CreateDate("sha256", sha256, date),
				*model.CreateDate("base64", base64, date),
			},
		)
		break
	default:
		err = errors.New("Invalid type of hash code: " + typeFile + "!\n")
		break
	}
	return err
}

/*
 Get data which has specified "date" field
*/
func GetData(date string, typeFile string) (results []model.Data, err error) {
	defer func() {
		if _err := recover(); _err != nil {
			results = []model.Data{}
			err = _err.(error)
		}
	}()

	results = []model.Data{}
	err = nil

	if !regexpDate.MatchString(date) {
		return nil, errors.New(`Incompatible date format. Must be "yyyy-MM-dd".\n`)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch typeFile {
	case "md5", "sha1", "sha256", "base64":
		var data model.Data
		err := config.DB.Collection("malshare").FindOne(ctx, model.Data{
			Type: typeFile,
			Date: date,
		}).Decode(&data)
		utils.CatchError(err)
		results = append(results, data)
		break

	case "all", "":
		cursor, err := config.DB.Collection("malshare").Find(ctx, model.Data{
			Date: date,
		})
		utils.CatchError(err)
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var data model.Data
			cursor.Decode(&data)
			results = append(results, data)
		}
		utils.CatchError(cursor.Err())
		break

	default:
		err = errors.New("Invalid type of hash code: " + typeFile + "!\n")
		break

	}

	return results, err
}
