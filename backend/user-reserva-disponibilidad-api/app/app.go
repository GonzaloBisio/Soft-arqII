package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"user-reserva-disponibilidad-api/cache"
	db2 "user-reserva-disponibilidad-api/db"
)

/*func init() {

	router = gin.Default()

	router.Static("/images", "./images")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.AddExposeHeaders("X-Total-Count")
	router.Use(cors.New(config))

	log.SetOutput(os.Stdout)
	//log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting logger system")
}*/

var router *gin.Engine

func StartApp() {
	router = gin.Default()
	mapUrls()

	db2.DatabaseCon()
	db2.StartDbEngine()
	router.Run(":8003")
	log.Info("Starting Server")

	cache.Init_cache()
	data := []byte(`{"key": "value"}`)
	cache.Set("mi_llave", data, 10)

	key := "mi_llave"
	availability, err := cache.Get(key)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			log.Println("La llave no se encontró en el caché")
		} else {
			log.Println("Error al recuperar datos del caché:", err)
		}
	} else {
		log.Println("Datos recuperados del caché:", availability)
	}
	go func() {
		time.Sleep(10 * time.Second)
	}()

	availability, err = cache.Get(key)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			log.Println("La llave no se encontró en el caché")
		} else {
			log.Println("Error al recuperar datos del caché:", err)
		}
	} else {
		log.Println("Datos recuperados del caché:", availability)
	}
}
