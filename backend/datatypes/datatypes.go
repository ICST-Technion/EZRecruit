package datatypes

import "github.com/ICST-Technion/EZRecruit.git/queries"

// JobListing represents a job-listing entry.
type JobListing struct {
	ID string `json:"_id" form:"_id"`
	queries.PostJobListing
}

// JobApplication represents a job-application entry.
type JobApplication struct {
	ID string `json:"_id" form:"_id"`
	queries.PostJobApplication
	ResumeFileLocation string
}
