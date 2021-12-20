package restapi

import (
	"fmt"
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
	"github.com/ICST-Technion/EZRecruit.git/pkg/db"
	"github.com/ICST-Technion/EZRecruit.git/queries"
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
	router.GET("/applications", s.getJobApplications)
	// POST METHODS
	router.POST("/jobs", s.insertJobListing)
	router.POST("/applications", s.insertJobApplication)
	// DELETE METHODS
	router.DELETE("/jobs", s.deleteJobListing)
}

/// #############################################
/// ####### JOB LISTING ENDPOINT HANDLERS #######
/// #############################################

// getJobListings responds with a list of the job-listings relevant to the request.
func (s *Server) getJobListings(c *gin.Context) {
	query := &queries.GetJobListing{
		Labels: []string{},
	}

	if err := c.ShouldBind(query); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, s.dbClient.GetJobs(query.Labels))
}

// insertJobListing inserts a job to the database.
func (s *Server) insertJobListing(c *gin.Context) {
	jobListing := &datatypes.JobListing{}

	if err := c.ShouldBind(jobListing); err != nil {
		fmt.Printf("failed to bind object in insertJobRequest - %v\n", err)
		return
	}

	fmt.Printf("adding job - %v\n", *jobListing)
	// Add jobListing to collection.
	id := s.dbClient.InsertJob(jobListing)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("job with id {%s} updated", id)})
}

// deleteJobListing deletes a job listing by ID from the database.
func (s *Server) deleteJobListing(c *gin.Context) {
	if id, found := c.GetQuery("id"); found && s.dbClient.DeleteJob(id) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "job deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}

/// #############################################
/// ##### JOB APPLICATION ENDPOINT HANDLERS #####
/// #############################################

// getJobApplications responds with a list of the job-applications relevant to the request.
func (s *Server) getJobApplications(c *gin.Context) {
	query := &queries.GetJobApplication{
		Labels: []string{},
	}

	if err := c.ShouldBind(query); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, s.dbClient.GetApplications(query.Labels))
}

// insertJobListing inserts a job application to the database.
func (s *Server) insertJobApplication(c *gin.Context) {
	// application info
	jobApplication := &datatypes.JobApplication{}
	if err := c.ShouldBind(jobApplication); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// resume file
	file, err := c.FormFile("file")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobApplication.ResumeFileLocation = fmt.Sprintf("%s_%s", jobApplication.User, file.Filename)

	// save file
	err = c.SaveUploadedFile(file, "resume/"+jobApplication.ResumeFileLocation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to save resume file, contact admin"})
		fmt.Printf("failed to save resume file %s - %v", jobApplication.ResumeFileLocation, err)

		return
	}

	fmt.Printf("adding application - %v\n", jobApplication)
	// Add jobListing to collection.
	id := s.dbClient.InsertApplication(jobApplication)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("recieved job application - %s", id)})
}
