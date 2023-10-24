package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoHost     = "mongodb://localhost:27017"
	mongoDatabase = "proyecto-arquiII"
)

var Client *MongoClient

type MongoClient struct {
	ImageCollection *mongo.Collection
	HotelCollection *mongo.Collection
}

func InitializeMongoClient() {
	clientOptions := options.Client().ApplyURI(mongoHost)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	HotelCollection := client.Database(mongoDatabase).Collection("hoteles")
	ImageCollection := client.Database(mongoDatabase).Collection("images")

	Client = &MongoClient{
		HotelCollection: HotelCollection,
		ImageCollection: ImageCollection,
	}
}
