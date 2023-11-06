package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	db2 "user-reserva-disponibilidad-api/db"
)

func init() {

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
}

var router *gin.Engine

func StartApp() {

	mapUrls()

	db2.DatabaseCon()
	db2.StartDbEngine()

	router := gin.Default()
	router.Run("localhost:8002")
	log.Info("Starting Server")

}
