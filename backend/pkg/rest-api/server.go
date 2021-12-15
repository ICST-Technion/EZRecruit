package restapi

import (
	"github.com/ICST-Technion/EZRecruit.git/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRestAPIServer returns a new instance of RestAPIServer.
func NewRestAPIServer(dbClient db.DB) *Server {
	return &Server{
		dbClient: dbClient,
	}
}

// Server implements the restAPI functionality logic
type Server struct {
	dbClient db.DB
}

// GetJobListings responds with the list of all job-listings as JSON.
func (s *Server) GetJobListings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, s.dbClient.GetJobs())
}
