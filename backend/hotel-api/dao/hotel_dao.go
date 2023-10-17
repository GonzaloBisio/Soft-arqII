package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	mg "hotel-api/app"
	"hotel-api/models"
	"log"
)

type HotelDAO struct {
	Collection *mongo.Collection
}

// InitializeHotelDAO creates a new HotelDAO with the specified collection name.
func (c *HotelDAO) InitializeHotelDAO(collectionName string) *HotelDAO {
	var collection *mongo.Collection

	if collectionName == "hoteles" {
		collection = mg.Client.HotelCollection
	} else if collectionName == "images" {
		collection = mg.Client.ImageCollection
	} else {
		log.Fatal("Invalid collection name")
	}

	if collection == nil {
		log.Fatal("MongoDB collection not initialized")
	}

	return &HotelDAO{
		Collection: collection,
	}
}

func (c *HotelDAO) GetAll() ([]models.Hotel, error) {
	var hotels []models.Hotel
	cursor, err := mg.Client.HotelCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var hotel models.Hotel
		err := cursor.Decode(&hotel)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (c *HotelDAO) Insert(hotel models.Hotel) (models.Hotel, error) {
	// Omit configuring the _id field
	request, err := mg.Client.HotelCollection.InsertOne(context.Background(), hotel)

	if err != nil {
		return hotel, err
	}

	// The insertion was successful, return the hotel with its updated ID
	insertedHotel := hotel
	insertedHotel.ID = request.InsertedID.(primitive.ObjectID)

	return insertedHotel, nil
}

func (c *HotelDAO) GetHotelById(id string) (models.Hotel, error) {
	var hotel models.Hotel
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return hotel, err
	}
	err = mg.Client.HotelCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&hotel)
	return hotel, err
}

func (c *HotelDAO) Update(hotel models.Hotel) (models.Hotel, error) {
	_, err := mg.Client.HotelCollection.ReplaceOne(context.Background(), bson.M{"_id": hotel.ID}, hotel)
	if err != nil {
		return models.Hotel{}, err
	}
	return hotel, nil
}
