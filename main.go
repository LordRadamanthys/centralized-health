package main

import (
	"github.com/LordRadamanthys/centralized-health/adapter/input/routes"
	"github.com/LordRadamanthys/centralized-health/configuration/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	mongodb.InitMongoDBConnection()

	routes.RoutesUrl()
}
