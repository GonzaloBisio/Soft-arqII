package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	db "search-api/db"

	dto "search-api/dto"

	"github.com/stevenferrer/solr-go"
)

var solrDB *db.Solr

func init() {
	solrDB, _ = db.NewSolrConnection()
}

func AddFromId(hotel dto.HotelDTO) (dto.HotelDTO, error) {
	// Preparo los dto con formato de archivo

	var addHotelDto dto.AddDto
	addHotelDto.Add = dto.DocDto{Doc: hotel}
	//CReo el json
	data, err := json.Marshal(addHotelDto)
	//Creo el reader
	reader := bytes.NewReader(data)

	// Realiza la actualización en Solr
	_, err = solrDB.Client.Update(context.TODO(), solrDB.Collection, solr.JSON, reader)

	if err != nil {

		return hotel, err
	}

	// Confirma los cambios en Solr
	err = solrDB.Client.Commit(context.TODO(), solrDB.Collection)

	if err != nil {
		log.Println("Error en el commit linea 45")

		return hotel, err
	}
	return hotel, nil
}

func DeleteFromId(id string) error {
	var deleteDto dto.DeleteDto
	deleteDto.Delete = dto.DeleteDoc{Query: fmt.Sprintf("id:%s", id)}
	data, err := json.Marshal(deleteDto)
	reader := bytes.NewReader(data)

	if err != nil {
		return err
	}
	_, err = solrDB.Client.Update(context.TODO(), solrDB.Collection, solr.JSON, reader)

	if err != nil {
		return err
	}

	err = solrDB.Client.Commit(context.TODO(), solrDB.Collection)
	if err != nil {
		return err
	}

	return nil
}
