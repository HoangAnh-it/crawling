package utils

import (
	"crawling/concurrency"
	"crawling/config"
	"crawling/model"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	urlPattern      = `href="(\d{4}-\d{2}-\d{2})/"`
	folderPattern   = `\d{4}-\d{2}-\d{2}`
	OUTPUT_PATH     = "output"
	hashCodePattern = `(\s?[a-f0-9]{32}\b)?(\s[0-9a-f]{40}\b)?(\s[A-Fa-f0-9]{64}\b)?(\s[a-zA-Z0-9:\+/]+\b)?`
)

var (
	regexpUrl     = regexp.MustCompile(urlPattern)
	regexpFolder  = regexp.MustCompile(folderPattern)
	regexHashCode = regexp.MustCompile(hashCodePattern)
)

func makeRequest(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fail: preparing request GET to %s.\n", url))
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fail: making GET to %s.\n", url))
	}
	return response, nil
}

func doRequest(url string) *http.Response {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	response, err := makeRequest(url)
	if err != nil {
		panic(err)
	}

	if response == nil {
		panic(fmt.Sprintln("No data in", url))
	}

	if response.StatusCode != 200 {
		panic(fmt.Sprintf("GET %s return status: %d\n", url, response.StatusCode))
	}

	fmt.Printf("Done: GET %s\n", url)
	return response
}

func Crawling(url string) {
	response := doRequest(url)
	defer response.Body.Close()
	responseText := ConvertToResponseText(response)
	links := regexpUrl.FindAllStringSubmatch(responseText, -1)

	for _, link := range links {
		date := link[1]
		folderPath := createFolder(date)

		concurrency.JobList <- concurrency.NewJob(func() {
			linkFileText := fmt.Sprintf("https://malshare.com/daily/%s/malshare_fileList.%s.all.txt", date, date)
			responseContent := doRequest(linkFileText)
			defer func() {
				if responseContent != nil {
					responseContent.Body.Close()
				}
			}()
			content := ConvertToResponseText(responseContent)
			saveToLocal(folderPath, date, content)
		})
	}
}

func saveToLocal(folderPath string, date string, content string) {
	lines := strings.Split(content, "\n")
	var md5Arr, sha1Arr, sha256Arr, base64Arr []string
	var data []model.Data
	for _, line := range lines {
		line = " " + strings.TrimSpace(line)
		matchString := regexHashCode.FindAllStringSubmatch(line, -1)[0]
		md5 := strings.TrimSpace(matchString[1])
		sha1 := strings.TrimSpace(matchString[2])
		sha256 := strings.TrimSpace(matchString[3])
		base64 := strings.TrimSpace(matchString[4])

		md5Arr = append(md5Arr, md5)
		sha1Arr = append(sha1Arr, sha1)
		sha256Arr = append(sha256Arr, sha256)
		base64Arr = append(base64Arr, base64)

		data = append(data, model.Data{
			Md5:    md5,
			Sha1:   sha1,
			Sha256: sha256,
			Base64: base64,
			Date:   date,
		})
	}

	createFile(strings.Join(md5Arr, " "), "md5.txt", folderPath)
	createFile(strings.Join(sha1Arr, " "), "sha1.txt", folderPath)
	createFile(strings.Join(sha256Arr, " "), "sha256.txt", folderPath)
	createFile(strings.Join(base64Arr, " "), "base64.txt", folderPath)
	go config.DBController.InsertMany(ConvertToInterface(data))
}

func createFolder(date string) string {
	folderName := OUTPUT_PATH + "/" + strings.ReplaceAll(date, "-", "/")

	_, err := os.Stat(folderName)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(folderName, os.ModePerm)
		fmt.Println("Create folder:", folderName)
	}

	currentPath, _ := os.Getwd()
	return filepath.Join(currentPath, folderName)
}

func createFile(content string, name string, folderPath string) {
	path := filepath.Join(folderPath, name)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Cannot create file: ", path)
		fmt.Println(err)
		return
	}
	defer file.Close()
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if _, err := file.WriteString(content); err != nil {
		fmt.Println("Cannot write to", path)
		return
	}
	fmt.Println("Save to", path)
}
