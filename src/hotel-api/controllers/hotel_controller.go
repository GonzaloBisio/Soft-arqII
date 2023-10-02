package controllers

import (
	"net/http"
)

func CreateHotel(w http.ResponseWriter, r *http.Request) {
    // Lógica para crear un hotel en MongoDB y notificar a RabbitMQ
}

func GetHotelByID(w http.ResponseWriter, r *http.Request) {
    // Lógica para obtener un hotel por ID
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
    // Lógica para modificar un hotel en MongoDB y notificar a RabbitMQ
}

// ... otros controladores
