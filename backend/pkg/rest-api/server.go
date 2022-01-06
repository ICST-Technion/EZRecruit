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
	envVarContainerPort  = "PORT"
	resumeFolderLocation = "/usr/local/bin/resume"
	workDir              = "/usr/local/bin"
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
	router.LoadHTMLFiles(fmt.Sprintf("%s/%s", workDir, "upload_resume.html"))
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_resume.html", gin.H{})
	})
	// GET METHODS
	router.GET("/jobs", s.getJobListings)
	router.GET("/applications", s.getJobApplications)
	router.GET("/resume", s.getApplicantResume)
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
		Filterable: &queries.Filterable{FilterLabels: []string{}},
		Sortable:   &queries.Sortable{SortLabels: []string{}},
		Pagination: &queries.Pagination{},
	}

	if err := c.ShouldBind(query); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	jobListings := s.dbClient.GetJobs(query.Filterable, query.Sortable)

	if query.Pagination.Limit == 0 {
		c.IndentedJSON(http.StatusOK, datatypes.PaginatedResponse{Size: len(jobListings), Value: jobListings})
	} else if query.Pagination.Limit >= len(jobListings) {
		c.IndentedJSON(http.StatusOK, jobListings) // result smaller than limit
	} else if (query.Pagination.Offset+1)*query.Pagination.Limit >= len(jobListings) {
		c.IndentedJSON(http.StatusOK, datatypes.PaginatedResponse{Size: len(jobListings),
			Value: jobListings[query.Pagination.Offset*query.Pagination.Limit:]}) // chunk left
	} else {
		c.IndentedJSON(http.StatusOK, datatypes.PaginatedResponse{Size: len(jobListings),
			Value: jobListings[query.Pagination.Offset*query.Pagination.Limit : (query.Pagination.Offset+1)*query.Pagination.Limit]}) // pagination
	}
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
func (s *Server) getApplicantResume(c *gin.Context) {
	if user, found := c.GetQuery("user"); found {
		if file, exists := s.dbClient.GetUserResumeFileLocation(user); exists {
			c.IndentedJSON(http.StatusOK, gin.H{"message": file})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user does not have a resume file saved"})
}

// getJobApplications responds with a list of the job-applications relevant to the request.
func (s *Server) getJobApplications(c *gin.Context) {
	query := &queries.GetJobApplication{
		Filterable: &queries.Filterable{FilterLabels: []string{}},
		Sortable:   &queries.Sortable{SortLabels: []string{}},
	}

	if err := c.ShouldBind(query); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, s.dbClient.GetApplications(query.Filterable, query.Sortable))
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

	fileSaveLocation := fmt.Sprintf("%s_%s", jobApplication.User, file.Filename)

	// save file
	err = c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", resumeFolderLocation, fileSaveLocation))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to save resume file, contact admin"})
		fmt.Printf("failed to save resume file %s - %v", fileSaveLocation, err)

		return
	}

	fmt.Printf("adding application - %v\n", jobApplication)
	// Add jobListing to collection.
	id := s.dbClient.InsertApplication(jobApplication, fileSaveLocation)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("recieved job application - %s", id)})
}
