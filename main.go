package main

import (
	"crawling/concurrency"
	"crawling/config"
	"crawling/utils"
)

func main() {
	config.Setup()
	config.DBController.ConnectToMongo()
	URL := "https://malshare.com/daily/"
	manager := concurrency.NewManager(100)
	manager.Do(func() {
		utils.Crawling(URL)
	})
	concurrency.WG.Wait()
	manager.Finish()
}
