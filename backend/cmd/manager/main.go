package main

import (
	"github.com/ICST-Technion/EZRecruit.git/pkg/db/in-memory"
	"github.com/ICST-Technion/EZRecruit.git/pkg/rest-api"
)

func main() {
	// in memory DB
	inMemoryDB := inmemory.NewInMemoryDB()
	// restAPI server
	restAPIServer := restapi.NewRESTAPIServer(inMemoryDB)
	// start server
	restAPIServer.Start()
}
