package main

import (
	app "search-api/app"
	"search-api/db"
	queue "search-api/utils/connection"
)

var c *db.Solr

func main() {
	go queue.QueueConnection()
	c, _ = db.NewSolrConnection()

	app.StartApp()

}
