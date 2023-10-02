package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Import the "gorilla/mux" package
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"hotel-api/dao"    // Replace with the correct import path to your dao package
	"hotel-api/models" // Replace with the correct import path to your models package
)

var db *mgo.Database

func init() {
	// Initialize the MongoDB database connection in the init function
	session, err := mgo.Dial("127.0.0.1:27023")
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB("proyecto-arquiII")
}

func AllHotelesEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var hotel models.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	hotel.ID = bson.NewObjectId()

	// Create an instance of HotelesDAO and then call the Insert method on it
	daoInstance := &dao.HotelesDAO{} // Create an instance of HotelesDAO
	if err := daoInstance.Insert(hotel); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, hotel)
}

func FindHotelEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "can't find this particular hotel!")
}

func CreateHotelEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var hotel models.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	hotel.ID = bson.NewObjectId()

	// Create an instance of HotelesDAO and then call the Insert method on it
	daoInstance := &dao.HotelesDAO{} // Create an instance of HotelesDAO
	daoInstance.Connect()            // Connect to the MongoDB database

	if err := daoInstance.Insert(hotel); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, hotel)
}
func main() {
	r := mux.NewRouter() // Create a new router using "gorilla/mux"
	r.HandleFunc("/hoteles", AllHotelesEndpoint).Methods("GET")
	r.HandleFunc("/hotel/{id}", FindHotelEndpoint).Methods("GET")
	r.HandleFunc("/hotel", CreateHotelEndpoint).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
