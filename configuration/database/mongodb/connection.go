package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoConnection *mongo.Database
)

const (
	MONGO_DB_URL          = "MONGO_DB_URL"
	MONGO_HEALTH_DATABASE = "MONGO_HEALTH_DATABASE"
)

func InitMongoDBConnection() {
	mongoUrl := os.Getenv(MONGO_DB_URL)
	mongoHealthDB := os.Getenv(MONGO_HEALTH_DATABASE)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))

	if err != nil {
		panic(err)
	}

	err = client.Connect(context.TODO())

	if err != nil {
		panic(err)
	}

	MongoConnection = client.Database(mongoHealthDB)

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
}
