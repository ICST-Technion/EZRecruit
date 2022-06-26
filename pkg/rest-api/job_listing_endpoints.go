package restapi

import (
	"fmt"
	"github.com/ICST-Technion/EZRecruit/datatypes"
	"github.com/ICST-Technion/EZRecruit/queries"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getJobListings responds with a list of the job-listings relevant to the request.
func (s *Server) getJobListings(ctx *gin.Context) {
	query := &queries.GetJobListing{
		Filterable: &queries.Filterable{FilterLabels: []string{}},
		Sortable:   &queries.Sortable{SortLabels: []string{}},
		Pagination: &queries.Pagination{},
	}

	if err := ctx.ShouldBind(query); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("failed to bind object in getJobListings - %v\n", err)

		return
	}

	jobListings := s.dbClient.GetJobs(query.Filterable, query.Sortable)

	if query.Pagination.Limit == 0 {
		ctx.IndentedJSON(http.StatusOK, datatypes.PaginatedResponse{Size: len(jobListings), Value: jobListings})
		return
	}

	if query.Pagination.Limit >= len(jobListings) { // result smaller than limit
		ctx.IndentedJSON(http.StatusOK, datatypes.PaginatedResponse{Size: len(jobListings), Value: jobListings})
		return
	}

	if (query.Pagination.Offset+1)*query.Pagination.Limit >= len(jobListings) {
		ctx.IndentedJSON(http.StatusOK, datatypes.PaginatedResponse{
			Size:  len(jobListings),
			Value: jobListings[query.Pagination.Offset*query.Pagination.Limit:],
		}) // chunk left

		return
	}

	ctx.IndentedJSON(http.StatusOK, datatypes.PaginatedResponse{
		Size: len(jobListings),
		Value: jobListings[query.Pagination.Offset*query.Pagination.Limit : (query.Pagination.Offset+1)*
			query.Pagination.Limit],
	}) // pagination
}

// insertJobListing inserts a job to the database.
func (s *Server) insertJobListing(ctx *gin.Context) {
	jobListing := &datatypes.JobListing{}

	if err := ctx.ShouldBind(jobListing); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("failed to bind object in insertJobRequest - %v\n", err)

		return
	}
	// Add jobListing to collection.
	id := s.dbClient.InsertJob(jobListing)

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("job with id {%s} updated", id)})
}

// deleteJobListing deletes a job listing by ID from the database.
func (s *Server) deleteJobListing(ctx *gin.Context) {
	if id, found := ctx.GetQuery("id"); found && s.dbClient.DeleteJob(id) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "job deleted"})
		return
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}
