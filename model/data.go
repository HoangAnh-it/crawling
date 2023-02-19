package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Type     string             `bson:"type,omitempty"`
	HashCode string             `bson:"hash_code,omitempty"`
	Date     string             `bson:"date,omitempty"`
}

func CreateData(_type string, hashCode string, date string) *Data {
	return &Data{
		Type:     _type,
		HashCode: hashCode,
		Date:     date,
	}
}
