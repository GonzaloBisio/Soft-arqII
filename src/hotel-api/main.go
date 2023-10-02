package main

import (
	"net/http"

	"../hotel-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // Rutas para hoteles
    router.HandleFunc("/hotels", controllers.CreateHotel).Methods("POST")
    router.HandleFunc("/hotels/{id}", controllers.GetHotelByID).Methods("GET")
    router.HandleFunc("/hotels/{id}", controllers.UpdateHotel).Methods("PUT")

    http.ListenAndServe(":8080", router)
}
