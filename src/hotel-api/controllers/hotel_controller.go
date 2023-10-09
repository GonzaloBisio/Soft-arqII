package controllers

import (
	"encoding/json"
	"net/http"

	"hotel-api/models"
	"hotel-api/services"

	"github.com/gorilla/mux"
)

func CreateHotel(w http.ResponseWriter, r *http.Request) {
    // Aquí puedes implementar la lógica para crear un nuevo hotel.
    // Puedes usar el cuerpo de la solicitud (r.Body) para obtener los datos del hotel
    // y luego utilizar el servicio HotelService para insertar el hotel en la base de datos.
    // Por ejemplo:

    var newHotel models.Hotel
    err := json.NewDecoder(r.Body).Decode(&newHotel)
    if err != nil {
        http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
        return
    }

    // Llama a la función InsertHotel del servicio HotelService
    createdHotel, apiErr := services.HotelService.InsertHotel(newHotel)
    if apiErr != nil {
        http.Error(w, apiErr.Message, apiErr.Status)
        return
    }

    // Devuelve el hotel creado en la respuesta
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdHotel)
}

func GetHotelByID(w http.ResponseWriter, r *http.Request) {
    // Aquí puedes implementar la lógica para obtener un hotel por su ID.
    // Puedes utilizar el paquete "github.com/gorilla/mux" para obtener el ID desde
    // la ruta de la solicitud (por ejemplo, r.URL.Path).

    vars := mux.Vars(r)
    hotelID := vars["id"]

    // Llama a la función GetHotelById del servicio HotelService
    hotel, apiErr := services.HotelService.GetHotelById(hotelID)
    if apiErr != nil {
        http.Error(w, apiErr.Message, apiErr.Status)
        return
    }

    // Devuelve el hotel en la respuesta
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(hotel)
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
    // Aquí puedes implementar la lógica para actualizar un hotel en la base de datos.
    // Similar a la función CreateHotel, debes obtener los datos del hotel del cuerpo
    // de la solicitud y luego utilizar el servicio HotelService para realizar la actualización.
}

// ... otros controladores
