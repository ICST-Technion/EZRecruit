package restapi

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ICST-Technion/EZRecruit/pkg/db"
	"github.com/gin-gonic/gin"
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
	router.POST("/status", s.updateApplicantsStatus)
	// DELETE METHODS
	router.DELETE("/jobs", s.deleteJobListing)
}
