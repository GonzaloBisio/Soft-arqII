package db

import (
	Clients "user-reserva-disponibilidad-api/clients"
	"user-reserva-disponibilidad-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func DatabaseCon() {
	//Parametros de coneccion
	DBNombre := "arqsoft2"
	DBUser := "usuario"   //root
	DBPass := "usuario"   //password
	DBHost := "localhost" //JULIAN: Tengo mysql nativo en ubuntu y no me deja lo de la maquina virtual
	//POr eso dejo esto asi para mi xd
	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBNombre+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	Clients.Db = db

}

func StartDbEngine() {

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Reservation{})

	log.Info("Llegue hasta aca pana")
}
