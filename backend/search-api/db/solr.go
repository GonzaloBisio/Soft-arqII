package db

import (
	"search-api/config"

	"github.com/stevenferrer/solr-go"
)

type Solr struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrConnection() (*Solr, error) {
	// Crea una nueva instancia de Solr
	client := solr.NewJSONClient(config.SolrURL)

	return &Solr{
		Client:     client,
		Collection: config.SolRCollection,
	}, nil
}
