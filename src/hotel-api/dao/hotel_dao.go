package dao

import (
	"context"
	"fmt"
	"hotel-api/models"
	"hotel-api/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelDAO struct{}

func (dao *HotelDAO) GetAll() ([]models.Hotel, error) {
    db := db.MongoDB
    var hotels []models.Hotel

    cursor, err := db.Collection("hotels").Find(context.TODO(), bson.D{})
    if err != nil {
        fmt.Println(err)
        return hotels, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var hotel models.Hotel
        if err := cursor.Decode(&hotel); err != nil {
            fmt.Println(err)
            return hotels, err
        }
        hotels = append(hotels, hotel)
    }

    if err := cursor.Err(); err != nil {
        fmt.Println(err)
        return hotels, err
    }

    return hotels, nil
}

func (dao *HotelDAO) GetHotelByID(id string) (models.Hotel, error) {
    db := db.MongoDB
    var hotel models.Hotel

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        fmt.Println(err)
        return hotel, err
    }

    err = db.Collection("hotels").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&hotel)
    if err != nil {
        fmt.Println(err)
        return hotel, err
    }

    return hotel, nil
}

func (dao *HotelDAO) InsertHotel(hotel models.Hotel) (models.Hotel, error) {
    db := db.MongoDB
    insertHotel := hotel
    insertHotel.ID = primitive.NewObjectID()

    _, err := db.Collection("hotels").InsertOne(context.TODO(), &insertHotel)
    if err != nil {
        fmt.Println(err)
        return hotel, err
    }

    return hotel, nil
}

// UpdateHotel actualiza un hotel.
func (dao *HotelDAO) UpdateHotel(hotel models.Hotel) (models.Hotel, error) {
    db := db.MongoDB
    filter := bson.M{"_id": hotel.ID}
    update := bson.M{
        "$set": bson.M{
            "name":        hotel.Name,
            "description": hotel.Description,
        },
    }

    _, err := db.Collection("hotels").UpdateOne(context.TODO(), filter, update)
    if err != nil {
        fmt.Println(err)
        return hotel, err
    }

    return hotel, nil
}
