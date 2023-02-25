package utils

import (
	"crawling/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func ConvertToResponseText(response *http.Response) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	responsePlainText := string(res)
	return responsePlainText
}

func ConvertToInterface(data []model.Data) []interface{} {
	results := make([]interface{}, 0, len(data))
	for _, d := range data {
		results = append(results, d)
	}
	return results
}

func GetCurrentDate() string {
	layout := "2006-01-02"
	return time.Now().UTC().Format(layout)
}

func ExtractQueries(request *http.Request) DataModel {
	results := DataModel{}
	queries := request.URL.Query()

	if queries.Get("md5") != "" {
		results.Md5 = queries.Get("md5")
	}
	if queries.Get("sha1") != "" {
		results.Sha1 = queries.Get("sha1")
	}
	if queries.Get("sha256") != "" {
		results.Sha256 = queries.Get("sha256")
	}
	if queries.Get("base64") != "" {
		results.Base64 = queries.Get("base64")
	}
	if queries.Get("date") != "" {
		results.Date = queries.Get("date")
	}
	return results
}
