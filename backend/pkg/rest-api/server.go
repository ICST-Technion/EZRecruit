package restapi

import (
	"fmt"
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
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
	// POST METHODS
	router.POST("/jobs", s.insertJobListing)
	// DELETE METHODS
	router.DELETE("/jobs:_id", s.deleteJobListing)
}

// getJobListings responds with the list of all job-listings as JSON.
func (s *Server) getJobListings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, s.dbClient.GetJobs())
}

// insertJobListing inserts a job to the database.
func (s *Server) insertJobListing(c *gin.Context) {
	var jobListing datatypes.JobListing

	// Call BindJSON to bind the received JSON to jobListing.
	if err := c.BindJSON(&jobListing); err != nil {
		fmt.Printf("failed to bind json in insertJobRequest - %v\n", err)
		return
	}

	fmt.Printf("adding job - %v\n", jobListing)
	// Add jobListing to collection.
	id := s.dbClient.InsertJob(&jobListing)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("job with id {%s} updated", id)})
}

// deleteJobListing deletes a job listing by ID from the database.
func (s *Server) deleteJobListing(c *gin.Context) {
	id := c.Param("_id")

	if id != "" && s.dbClient.DeleteJob(id) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "job deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}
