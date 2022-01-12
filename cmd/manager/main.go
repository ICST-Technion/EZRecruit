package main

import (
	inmemory "github.com/ICST-Technion/EZRecruit/pkg/db/in-memory"
	restapi "github.com/ICST-Technion/EZRecruit/pkg/rest-api"
)

func main() {
	// in memory DB
	inMemoryDB := inmemory.NewInMemoryDB()
	// restAPI server
	restAPIServer := restapi.NewRESTAPIServer(inMemoryDB)
	// start server
	restAPIServer.Start()
}
