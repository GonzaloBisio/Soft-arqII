package dao

import (
	"context"
	"fmt"
	"hotel-api/models"
	_ "hotel-api/models"

	"gopkg.in/mgo.v2/bson"
)

/*
type HotelesDAO struct {
	Server   string
	Database string
}
*/

var (
	Client HotelClientInterface
)

type Client struct {}

type HotelClientInterface interface {
	GetAll() (models.Hotel, error)
	GetById(id string) (models.Hotel, error)
	Insert(hotel models.Hotel) (models.Hotel, error)
	Update(hotel models.Hotel) (models.Hotel, error)
}

/*
var db *mgo.Database

const (
	COLLECTION = "hoteles"
)

func (m *HotelesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
*/

func (HotelClientInterface Client) GetAll() (models.Hotel, error) {
	db := db.MongoDb 
	var hotels models.Hotels

	cursor, err := db.Collection("Hotels").Find(context.TODO(), bson.D{})
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

func (HotelClientInterface Client) GetById(id string) (models.Hotel, error) {
	var hotel models.Hotel
	db := db.MongoDb 

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return hotel, err
	}

	err = db.Collection("Hotels").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&hotel)
	if err != nil {
		fmt.Println(err)
		return hotel, err
	}

	return hotel, nil 

}