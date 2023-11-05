package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
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

	//DATABASE CONNECTION
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfull conection to MySql!")

	router := gin.Default()
	router.Run("localhost:8002")
	log.Info("Starting Server")

}
