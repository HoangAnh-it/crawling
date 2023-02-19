package main

import (
	"crawling/concurrency"
	"crawling/config"
	"crawling/routes"
	"crawling/utils"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	URL := "https://malshare.com/daily/"
	manager := concurrency.NewManager(100)
	manager.Do(func() {
		utils.Crawling(URL)
	})
	concurrency.WG.Wait()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Loading environment failed. Err: %s", err)
	}

	PORT := os.Getenv("PORT")
	HOST := os.Getenv("HOST")

	config.ConnectToMongo()
	server := &http.Server{
		Handler:      routes.InitRoutes("http", HOST, PORT),
		Addr:         fmt.Sprintf("%s:%s", HOST, PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on port:", PORT)
	log.Fatalln(server.ListenAndServe())
}
