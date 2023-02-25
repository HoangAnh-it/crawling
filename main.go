package main

import (
	"crawling/concurrency"
	"crawling/config"
	"crawling/routes"
	"crawling/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	config.Setup()
	config.DBController.ConnectToMongo()
	Crawl()
	StartServer()
}

func Crawl() {
	URL := "https://malshare.com/daily/"
	manager := concurrency.NewManager(100)
	manager.Do(func() {
		utils.Crawling(URL)
	})
	concurrency.WG.Wait()
	manager.Finish()
}

func StartServer() {
	PORT := os.Getenv("PORT")
	HOST := os.Getenv("HOST")
	server := &http.Server{
		Handler:      routes.InitRoutes("http", HOST, PORT),
		Addr:         fmt.Sprintf("%s:%s", HOST, PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on %s:%s...\n", HOST, PORT)
	log.Fatalln(server.ListenAndServe())
}
