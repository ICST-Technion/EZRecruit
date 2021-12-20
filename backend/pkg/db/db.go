package db

import "github.com/ICST-Technion/EZRecruit.git/datatypes"

// DB abstracts the functionality needed from a DB client.
type DB interface {
	JobListingsHandler
	JobApplicationsHandler
}

// JobListingsHandler abstracts the functionality needed from a DB client for handling job-listings.
type JobListingsHandler interface {
	// GetJobs function to return stored jobs. If labels is empty then returns all.
	GetJobs(labels []string) []datatypes.JobListing
	// InsertJob function to insert a job to the database. If the ID field is left empty, a unique ID is assigned.
	// Otherwise, the given job-listing overwrites the existing one.
	//
	// Returns the listing's ID.
	InsertJob(jobListing *datatypes.JobListing) string
	// DeleteJob function to delete a job given its unique identifier. If a relevant job is found and is deleted,
	// returns true. Otherwise, returns false.
	DeleteJob(jobId string) bool
}

// JobApplicationsHandler abstracts the functionality needed from a DB client for handling job-applications.
type JobApplicationsHandler interface {
	// InsertApplication function to insert a job application. If the ID field is left empty, a unique ID is assigned.
	// Otherwise, the given job-listing overwrites the existing one.
	//
	// Returns the listing's ID.
	InsertApplication(jobApplication *datatypes.JobApplication) string
	// GetApplications function to return stored job applications. If labels is empty then returns all.
	GetApplications(labels []string) []datatypes.JobApplication
}
