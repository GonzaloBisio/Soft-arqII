package app

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

func init() {

	router = gin.Default()
	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")
}

func StartApp() {

	// Inicializa el cliente MongoDB

	mapUrls()

	log.Info("Starting Server")
	router.Run(":8001")

}
