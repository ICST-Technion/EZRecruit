package main

import (
	"fmt"
	"github.com/ICST-Technion/EZRecruit.git/pkg/db/in-memory"
	"github.com/ICST-Technion/EZRecruit.git/pkg/rest-api"
	"github.com/gin-gonic/gin"
)

func main() {
	// in memory DB
	inMemoryDB := inmemory.NewInMemoryDB()
	// restAPI server
	restAPIServer := restapi.NewRestAPIServer(inMemoryDB)

	// the restAPI public methods
	router := gin.Default()
	router.GET("/jobs", restAPIServer.GetJobListings)

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Printf("failed to start server - %v", err)
		return
	}
}
