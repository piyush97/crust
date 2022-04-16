package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	route "github.com/piyush97/crust/routes"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title           Crust API
// @version         1.0
// @description     Crust Api Server.

// @contact.name   Piyush Mehta

// @license.name  Apache 2.0

// @host      localhost:8090
// @BasePath  /api

func main() {
	fmt.Println("Main Application Starts")
	//Loading Environmental Variable
	loadenv()
	log.Fatal(route.RunAPI(":8090"))

}
