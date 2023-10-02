package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"hotel-api/models"
	_ "hotel-api/models"
	"log"
)

type HotelesDAO struct {
	Server   string
	Database string
}

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

func (m *HotelesDAO) FindAll() ([]models.Hotels, error) {
	var movies []models.Hotels
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *HotelesDAO) FindById(id string) (models.Hotel, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *HotelesDAO) Insert(movie models.Hotel) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *HotelesDAO) Delete(movie models.Hotel) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

func (m *HotelesDAO) Update(movie models.Hotel) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
