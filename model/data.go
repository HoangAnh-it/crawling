package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Md5    string             `bson:"md5,omitempty" json:"md5"`
	Sha1   string             `bson:"sha1,omitempty" json:"sha1"`
	Sha256 string             `bson:"sha256,omitempty" json:"sha256"`
	Base64 string             `bson:"base64,omitempty" json:"base64"`
	Date   string             `bson:"date,omitempty" json:"date"`
}
