package utils

import (
	"errors"
)

func ValidateData(data *DataModel, restrict bool) {
	if data.Md5 == "" && data.Sha1 == "" && data.Sha256 == "" && data.Base64 == "" {
		panic(errors.New("Missing hash code value. Must have at least 1 hash code."))
	}

	if restrict && data.Date == "" {
		data.Date = GetCurrentDate()
	}
}
