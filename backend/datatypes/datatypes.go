package datatypes

import "github.com/ICST-Technion/EZRecruit/queries"

// JobListing represents a job-listing entry.
type JobListing struct {
	ID string `json:"_id" form:"_id"`
	queries.PostJobListing
}

// JobApplication represents a job-application entry.
type JobApplication struct {
	ID       string `json:"_id" form:"_id"`
	JobTitle string `json:"jobTitle" form:"jobTitle"`
	queries.PostJobApplication
}

// PaginatedResponse represents a paginated response.
type PaginatedResponse struct {
	Size  int         `form:"size" json:"size"`
	Value interface{} `form:"value" json:"value"`
}
