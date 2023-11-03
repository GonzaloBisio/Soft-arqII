package main

import (
	app "search-api/app"
	queue "search-api/utils/connection"
)

func main() {
	go queue.QueueConnection()
	app.StartApp()

}
