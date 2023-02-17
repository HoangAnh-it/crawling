package utils

import (
	"crawling/concurrency"
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
	urlPattern    = `href="(\d{4}-\d{2}-\d{2})/"`
	folderPattern = `\d{4}-\d{2}-\d{2}`
	OUTPUT_PATH   = "output"
)

var (
	regexpUrl    = regexp.MustCompile(urlPattern)
	regexpFolder = regexp.MustCompile(folderPattern)
)

func getType(url string) string {
	if strings.HasSuffix(url, ".all.txt") {
		return "file"
	}
	return "folder"
}

func makeRequest(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 15,
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

	if response.StatusCode != 200 {
		panic(fmt.Sprintf("GET %s return status: %d\n", url, response.StatusCode))
	}

	fmt.Printf("Done: GET %s\n", url)
	return response
}

func Crawling(url string) {
	response := doRequest(url)
	defer response.Body.Close()
	responseText := convertToResponseText(response)
	links := regexpUrl.FindAllStringSubmatch(responseText, -1)

	for _, link := range links {
		concurrency.JobList <- concurrency.NewJob(func() {
			concurrency.WG.Add(1)
			date := link[1]
			linkFileText := fmt.Sprintf("https://malshare.com/daily/%s/malshare_fileList.%s.all.txt", date, date)
			responseContent := doRequest(linkFileText)
			content := convertToResponseText(responseContent)
			saveToLocal(date, content)
			concurrency.WG.Done()
		})
	}
}

func saveToLocal(date string, content string) {
	contentFile := strings.Fields(content)
	var md5, sha1, sha256, base64 []string
	length := len(contentFile)
	for i := 0; i+3 < length; i = i + 4 {
		md5 = append(md5, contentFile[i])
		sha1 = append(sha1, contentFile[i+1])
		sha256 = append(sha256, contentFile[i+2])
		base64 = append(base64, contentFile[i+3])
	}
	folderPath := createFolder(date)
	createFile(strings.Join(md5, "\t"), "md5.txt", folderPath)
	createFile(strings.Join(sha1, "\t"), "sha1.txt", folderPath)
	createFile(strings.Join(sha256, "\t"), "sha256.txt", folderPath)
	createFile(strings.Join(base64, "\t"), "base64.txt", folderPath)
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
	defer file.Close()
	if err != nil {
		fmt.Println("Cannot create file: ", path)
		fmt.Println(err)
		return
	}
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
