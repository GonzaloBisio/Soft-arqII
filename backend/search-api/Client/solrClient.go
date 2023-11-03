package client

import (
	"bytes"
	"context"
	"encoding/json"
	db "search-api/db"

	dto "search-api/dto"

	"github.com/stevenferrer/solr-go"
)

func AddFromId(hotel dto.HotelDto) (dto.HotelDto, error) {
	// Crea una conexión Solr
	conection, err := db.NewSolrConnection("hoteles")

	if err != nil {
		return hotel, err
	}

	// Serializa el objeto AddDto en formato JSON
	data, err := json.Marshal(hotel)
	if err != nil {
		return hotel, err
	}

	// Crea un lector de bytes con los datos serializados
	reader := bytes.NewReader(data)

	// Realiza la actualización en Solr
	_, err = conection.Client.Update(context.TODO(), conection.Collection, solr.JSON, reader)
	if err != nil {
		return hotel, err
	}

	// Confirma los cambios en Solr
	err = conection.Client.Commit(context.TODO(), conection.Collection)
	if err != nil {
		return hotel, err
	}

	return hotel, nil
}
