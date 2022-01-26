package restapi

import (
	"fmt"
	"github.com/ICST-Technion/EZRecruit/datatypes"
	"github.com/ICST-Technion/EZRecruit/queries"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getJobApplications responds with a list of the job-applications relevant to the request.
func (s *Server) getApplicantResume(ctx *gin.Context) {
	if user, found := ctx.GetQuery("user"); found {
		if file, exists := s.dbClient.GetUserResumeFileLocation(user); exists {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": file})
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user does not have a resume file saved"})
}

// getJobApplications responds with a list of the job-applications relevant to the request.
func (s *Server) getJobApplications(ctx *gin.Context) {
	query := &queries.GetJobApplication{
		Filterable: &queries.Filterable{FilterLabels: []string{}},
		Sortable:   &queries.Sortable{SortLabels: []string{}},
	}

	if err := ctx.ShouldBind(query); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("failed to bind object in getJobApplications - %v\n", err)

		return
	}

	ctx.IndentedJSON(http.StatusOK, s.dbClient.GetApplications(query.Filterable, query.Sortable))
}

// insertJobListing inserts a job application to the database.
func (s *Server) insertJobApplication(ctx *gin.Context) {
	// application info
	jobApplication := &datatypes.JobApplication{}
	if err := ctx.ShouldBind(jobApplication); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// resume file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("failed to bind object in insertJobApplication - %v\n", err)

		return
	}

	fileSaveLocation := fmt.Sprintf("%s_%s", jobApplication.User, file.Filename)

	// save file
	err = ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", resumeFolderLocation, fileSaveLocation))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to save resume file, contact admin"})
		fmt.Printf("failed to save resume file %s - %v", fileSaveLocation, err)

		return
	}

	fmt.Printf("adding application - %v\n", jobApplication)
	// Add jobListing to collection.
	id := s.dbClient.InsertApplication(jobApplication, fileSaveLocation)

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("received job application - %s", id)})
}

// updateApplicantsStatus updates the application status for applicants in the database.
func (s *Server) updateApplicantsStatus(ctx *gin.Context) {
	updateApplicantsStatus := &queries.UpdateApplicantsStatus{}

	if err := ctx.ShouldBind(updateApplicantsStatus); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("failed to bind object in updateApplicantsStatus - %v\n", err)

		return
	}

	s.dbClient.SetApplicantsStatus(updateApplicantsStatus.Users, updateApplicantsStatus.StatusID)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "updated applicants"})
}
