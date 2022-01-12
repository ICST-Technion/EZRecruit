package db

import (
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
	"github.com/ICST-Technion/EZRecruit.git/queries"
)

// DB abstracts the functionality needed from a DB client.
type DB interface {
	JobListingsHandler
	JobApplicationsHandler
	ResumeFileHandler
}

// JobListingsHandler abstracts the functionality needed from a DB client for handling job-listings.
type JobListingsHandler interface {
	// GetJobs function to return stored jobs. If labels is empty then returns all.
	GetJobs(filterable *queries.Filterable, sortable *queries.Sortable) []datatypes.JobListing
	// InsertJob function to insert a job to the database. If the ID field is left empty, a unique ID is assigned.
	// Otherwise, the given job-listing overwrites the existing one.
	//
	// Returns the listing's ID.
	InsertJob(jobListing *datatypes.JobListing) string
	// DeleteJob function to delete a job given its unique identifier. If a relevant job is found and is deleted,
	// returns true. Otherwise, returns false.
	DeleteJob(jobID string) bool
}

// JobApplicationsHandler abstracts the functionality needed from a DB client for handling job-applications.
type JobApplicationsHandler interface {
	// InsertApplication function to insert a job application. If the ID field is left empty, a unique ID is assigned.
	// Otherwise, the given job-listing overwrites the existing one.
	//
	// Returns the listing's ID.
	InsertApplication(jobApplication *datatypes.JobApplication, resumeLocation string) string
	// GetApplications function to return stored job applications. If labels is empty then returns all.
	GetApplications(filterable *queries.Filterable, sortable *queries.Sortable) []datatypes.JobApplication
}

// ResumeFileHandler abstracts the functionality needed from a DB client for handling resume files.
type ResumeFileHandler interface {
	// GetUserResumeFileLocation returns user's resume file location and whether it was found.
	GetUserResumeFileLocation(user string) (string, bool)
}
