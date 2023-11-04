package client

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	db "search-api/db"

	dto "search-api/dto"
	"strings"

	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

func AddFromId(hotel dto.HotelDTO) (dto.HotelDTO, error) {
	// Crea una conexión Solr

	conection, err := db.NewSolrConnection()

	var addHotelDto dto.AddDto
	addHotelDto.Add = dto.DocDto{Doc: hotel}
	data, err := json.Marshal(hotel)

	reader := bytes.NewReader(data)
	log.Println(addHotelDto)

	// Realiza la actualización en Solr
	resp, err := conection.Client.Update(context.TODO(), conection.Collection, solr.JSON, reader)
	logger.Debug(resp)
	if err != nil {
		log.Println(conection.Collection)

		return hotel, err
	}

	// Confirma los cambios en Solr
	err = conection.Client.Commit(context.TODO(), conection.Collection)
	if err != nil {
		log.Println("Error en el commit linea 45")

		return hotel, err
	}

	return hotel, nil
}

func Add(hotel dto.HotelDTO) error {
	client := solr.NewJSONClient("http://localhost:8983/solr/Hotels/")
	var addHotelDto dto.AddDto
	addHotelDto.Add = dto.DocDto{Doc: hotel}
	jsonData := `{"add":{"doc":` + `{"id":"` + hotel.ID + `","name":"` + hotel.Name + `","description":"` + hotel.Description + `"}}}`

	data := strings.NewReader(jsonData)

	// Realiza la actualización en Solr
	_, err := client.Update(context.TODO(), "mi_coleccion", solr.JSON, data)
	if err != nil {
		log.Fatalf("Error al agregar el objeto a Solr: %v", err)
	}
	return nil
}
