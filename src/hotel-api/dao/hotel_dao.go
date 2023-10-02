package dao

import (
	"gopkg.in/mgo.v2"
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
