package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"hotel-api/dao"
	"hotel-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Import the "gorilla/mux" package
)

func (m *dao.HotelesDAO) Insert(hotel models.Hotel) error {
	err := db.C(dao.COLLECTION).Insert(&hotel)
	return err
}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
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
	if err := dao.Insert(hotel); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, hotel)
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "can't find this particular hotel!")
}

func main() {
	r := mux.NewRouter() // Create a new router using "gorilla/mux"
	r.HandleFunc("/hoteles", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/hotel/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
