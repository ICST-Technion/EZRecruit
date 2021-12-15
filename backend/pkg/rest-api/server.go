package restapi

import (
	"fmt"
	"github.com/ICST-Technion/EZRecruit.git/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const (
	envVarContainerPort = "PORT"
)

// NewRESTAPIServer returns a new instance of RestAPIServer.
func NewRESTAPIServer(dbClient db.DB) *Server {
	return &Server{
		dbClient: dbClient,
	}
}

// Server implements the restAPI functionality logic.
type Server struct {
	dbClient db.DB
}

// Start function to start the listener.
func (s *Server) Start() {
	// the restAPI public methods
	router := gin.Default()
	// register handlers
	s.registerAPI(router)
	// Get port
	port := os.Getenv(envVarContainerPort)
	if port == "" {
		port = "8080"
	}

	// start
	fmt.Printf("started server - url:%s port:%s", os.Getenv("PUBLIC_URL"), port)

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Panicf("failed to start server - %v", err)
		return
	}
}

// registerAPI registers the handlers to the HTTP requests in router.
func (s *Server) registerAPI(router *gin.Engine) {
	// GET METHODS
	router.GET("/jobs", s.getJobListings)
}

// getJobListings responds with the list of all job-listings as JSON.
func (s *Server) getJobListings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, s.dbClient.GetJobs())
}
