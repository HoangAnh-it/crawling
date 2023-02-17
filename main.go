package main

import (
	"crawling/concurrency"
	"crawling/utils"
)

const (
	OUT_PUT_PATH = "output"
)

func main() {
	URL := "https://malshare.com/daily/"
	manager := concurrency.NewManager(100)
	manager.Do(func() {
		utils.Crawling(URL)
	})

	concurrency.WG.Wait()
}
