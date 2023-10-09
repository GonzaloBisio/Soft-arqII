package dao

import (
	"context"
	"hotel-api/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB configuration
var (
    mongoHost     = "mongodb://localhost:27017" 
    mongoDatabase = "proyecto-arquiII"  
)

// MongoDB client instance
var Client *MongoClient

// MongoClient represents the MongoDB client.
type MongoClient struct {
    Collection *mongo.Collection
}

// InitializeMongoClient initializes the MongoDB client.
func InitializeMongoClient() {
    clientOptions := options.Client().ApplyURI(mongoHost)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database(mongoDatabase).Collection("hoteles")
    Client = &MongoClient{
        Collection: collection,
    }
}

// GetAll retrieves all hotels from MongoDB.
func (c *MongoClient) GetAll() ([]models.Hotel, error) {
    var hotels []models.Hotel
    cursor, err := c.Collection.Find(context.Background(), bson.M{})
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

// Insert inserts a new hotel into MongoDB.
func (c *MongoClient) Insert(hotel models.Hotel) error {
    _, err := c.Collection.InsertOne(context.Background(), hotel)
    return err
}

// GetHotelById retrieves a hotel by its ID from MongoDB.
func (c *MongoClient) GetHotelById(id string) (models.Hotel, error) {
    var hotel models.Hotel
    var err error
    //objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return hotel, err
    }
    //err = c.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&hotel)
    return hotel, err
}

// Update updates a hotel in MongoDB.
func (c *MongoClient) Update(hotel models.Hotel) error {
    //objID, err := primitive.ObjectIDFromHex(hotel.ID)
    var err error
    if err != nil {
        return err
    }
    //_, err = c.Collection.ReplaceOne(context.Background(), bson.M{"_id": objID}, hotel)
    return err
}
