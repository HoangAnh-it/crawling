package utils

import "crawling/model"

type RawStringDataModel struct {
	Str  string `json:"str"`
	Date string `json:"date"`
}

type DataModel struct {
	Md5    string `json:"md5,omitempty" bson:"md5,omitempty"`
	Sha1   string `json:"sha1,omitempty" bson:"sha1,omitempty"`
	Sha256 string `json:"sha256,omitempty" bson:"sha256,omitempty"`
	Base64 string `json:"base64,omitempty" bson:"base64,omitempty"`
	Date   string `json:"date,omitempty" bson:"date,omitempty"`
}

func ToData(data DataModel) model.Data {
	return model.Data{
		Md5:    data.Md5,
		Sha1:   data.Sha1,
		Sha256: data.Sha256,
		Base64: data.Base64,
		Date:   data.Date,
	}
}
