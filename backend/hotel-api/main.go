package main

import (
	"hotel-api/dao"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
    dao.InitializeMongoClient()
}

func main() {
    r := mux.NewRouter()

    if err := http.ListenAndServe(":3000", r); err != nil {
        log.Fatal(err)
    }
}
