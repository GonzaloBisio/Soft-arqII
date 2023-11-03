package controllers

import (
	"log"

	"hotel-api/dto"
	"hotel-api/utils/queue"

	"github.com/gin-gonic/gin"
)

var RabbitMQConfig = queue.RabbitMQConfig{
	Username: "user",
	Password: "password",
	Host:     "localhost",
	Port:     "5672",
}

func PublishMessagePost(c *gin.Context) {
	c.Params.ByName("id")
	var qDto dto.QueueDto

	log.Println("Mandar mensaje")
	q, err := queue.NewRabbitMQQueue(RabbitMQConfig)

	if err != nil {
		log.Fatal("gol")
	}
	qDto.Action = "PUBLISH"
	qDto.Id = c.Params.ByName("id")
	q.PublishMessage(qDto)

}

func PublishMessageUpdate(c *gin.Context) {
	c.Params.ByName("id")
	var qDto dto.QueueDto

	log.Println("Mandar mensaje")
	q, err := queue.NewRabbitMQQueue(RabbitMQConfig)

	if err != nil {
		log.Fatal("gol")
	}
	qDto.Action = "UPDATE"
	qDto.Id = c.Params.ByName("id")
	q.PublishMessage(qDto)

}

func PublishMessageDelete(c *gin.Context) {
	c.Params.ByName("id")
	var qDto dto.QueueDto

	log.Println("Mandar mensaje")
	q, err := queue.NewRabbitMQQueue(RabbitMQConfig)

	if err != nil {
		log.Fatal("gol")
	}
	qDto.Action = "DELETE"
	qDto.Id = c.Params.ByName("id")
	q.PublishMessage(qDto)

}
