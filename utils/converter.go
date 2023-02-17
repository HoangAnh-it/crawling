package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func convertToResponseText(response *http.Response) string {
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
