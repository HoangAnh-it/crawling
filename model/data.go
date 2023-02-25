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

func (d *Data) Validate() bool {
	if d.Date == "" {
		return false
	}

	if d.Md5 != "" || d.Sha1 != "" || d.Sha256 != "" || d.Base64 != "" {
		return true
	}

	return false
}

func (d *Data) IsSame(other Data) bool {
	if d.Date != other.Date {
		return false
	}
	return d.Md5 == other.Md5 || d.Sha1 == other.Sha1 || d.Sha256 == other.Sha256 || d.Base64 == other.Base64
}
