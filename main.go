package main

import (
	"log"
	"os"

	"github.com/ShahSau/EthnicElegance/database"
	"github.com/ShahSau/EthnicElegance/router"
	"github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		log.Println("Loading .env file")
		err := godotenv.Load()

		if err != nil {
			log.Println("Error loading .env file")
		}
		log.Println("Loaded .env file successfully")
	}
	database.ConnectDB()
}

func main() {
	router.ClientRoutes() //run the routes
}
