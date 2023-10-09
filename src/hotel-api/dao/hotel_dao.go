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

type ProductionClient struct {}

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

func (HotelClientInterface ProductionClient) GetAll() (models.Hotel, error) {
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

func (HotelClientInterface ProductionClient) GetById(id string) (models.Hotel, error) {
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

func (HotelClientInterface ProductionClient) Insert(hotel models.Hotel) (models.Hotel, error) {
    db := db.MongoDb

    // Aquí debes implementar la lógica para insertar un nuevo hotel en la base de datos.
    // Puedes utilizar la función "InsertOne" del paquete "go.mongodb.org/mongo-driver/mongo".
    // Por ejemplo:

    insertResult, err := db.Collection("Hotels").InsertOne(context.TODO(), hotel)
    if err != nil {
        fmt.Println(err)
        return models.Hotel{}, err
    }

    // Si la inserción fue exitosa, puedes devolver el hotel insertado con su ID asignado
    // (puedes obtener el ID desde "insertResult.InsertedID").

    hotel.ID = insertResult.InsertedID.(primitive.ObjectID)
    return hotel, nil
}

func (HotelClientInterface ProductionClient) Update(hotel models.Hotel) (models.Hotel, error) {
    db := db.MongoDb

    // Aquí debes implementar la lógica para actualizar un hotel en la base de datos.
    // Puedes utilizar la función "UpdateOne" del paquete "go.mongodb.org/mongo-driver/mongo".
    // Por ejemplo:

    filter := bson.D{{"_id", hotel.ID}}
    update := bson.D{{"$set", bson.D{{"name", hotel.Name}, {"description", hotel.Description}}}}

    _, err := db.Collection("Hotels").UpdateOne(context.TODO(), filter, update)
    if err != nil {
        fmt.Println(err)
        return models.Hotel{}, err
    }

    return hotel, nil
}
