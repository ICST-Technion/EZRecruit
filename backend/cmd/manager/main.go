package main

import (
	"github.com/ICST-Technion/EZRecruit.git/pkg/db/in-memory"
	"github.com/ICST-Technion/EZRecruit.git/pkg/rest-api"
)

const envVarWorkDir = "WORK_DIR"

func main() {
	// in memory DB
	inMemoryDB := inmemory.NewInMemoryDB()
	// restAPI server
	restAPIServer := restapi.NewRESTAPIServer(inMemoryDB)
	// start server
	restAPIServer.Start()
}
