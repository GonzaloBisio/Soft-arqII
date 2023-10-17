package app

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

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
	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")
}

func StartApp() {

	mapUrls()

	log.Info("Starting Server")
	router.Run(":8000")

}
