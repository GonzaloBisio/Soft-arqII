package dao

import (
	"context"
	db "hotel-api/db"
	"hotel-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAll() ([]models.Hotel, error) {
	var hotels []models.Hotel
	cursor, err := db.Client.HotelCollection.Find(context.Background(), bson.M{})
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

func Insert(hotel models.Hotel) (models.Hotel, error) {

	request, err := db.Client.HotelCollection.InsertOne(context.Background(), hotel)

	if err != nil {
		return hotel, err
	}

	// The insertion was successful, return the hotel with its updated ID
	insertedHotel := hotel
	insertedHotel.ID = request.InsertedID.(primitive.ObjectID)

	return insertedHotel, nil
}

func GetHotelById(objID primitive.ObjectID) (models.Hotel, error) {
	var hotel models.Hotel

	err := db.Client.HotelCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&hotel)

	return hotel, err
}

func Update(hotel models.Hotel) (models.Hotel, error) {
	_, err := db.Client.HotelCollection.ReplaceOne(context.Background(), bson.M{"_id": hotel.ID}, hotel)
	if err != nil {
		return models.Hotel{}, err
	}
	return hotel, nil
}

func DeleteHotelById(id primitive.ObjectID) error {

	response, err := db.Client.HotelCollection.DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {
		return err
	}
	if response.DeletedCount == 0 {

		return nil // ACA HAY QUE CREAR UN ERROR DE QUE NO SE ENCONTRO HOTEL @AgusGlaiel
	}

	return nil
}
