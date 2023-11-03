package client

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	db "search-api/db"

	dto "search-api/dto"

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

func Add(hotelDto dto.HotelDTO) error {
	sc, _ := db.NewSolrConnection()
	var addHotelDto dto.AddDto
	addHotelDto.Add = dto.DocDto{Doc: hotelDto}
	data, err := json.Marshal(addHotelDto)
	log.Println(string(data))
	reader := bytes.NewReader(data)
	if err != nil {
		return err
	}
	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp)
	if err != nil {
		return err
	}

	er := sc.Client.Commit(context.TODO(), sc.Collection)
	if er != nil {
		logger.Debug("Error committing load")
		return err
	}
	return nil
}
