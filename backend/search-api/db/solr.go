package db

import (
	"fmt"
	"search-api/config"

	"github.com/stevenferrer/solr-go"
)

type Solr struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrConnection(collection string) (*Solr, error) {
	// Crea una nueva instancia de Solr
	client := solr.NewJSONClient(fmt.Sprintf("http://%s:%d/solr", config.SolRhost, config.SolRPort))

	return &Solr{
		Client:     client,
		Collection: collection,
	}, nil
}
