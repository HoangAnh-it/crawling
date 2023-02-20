package utils

import (
	"crawling/model"
	"fmt"
	"io/ioutil"
	"net/http"
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
