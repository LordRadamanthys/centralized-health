package main

import (
	"fmt"

	"github.com/LordRadamanthys/centralized-health/configuration/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	mongodb.InitMongoDBConnection()
	fmt.Println("teste")
}
