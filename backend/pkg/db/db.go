package db

import "github.com/ICST-Technion/EZRecruit.git/datatypes"

type DB interface {
	// GetJobs function to return stored jobs.
	GetJobs() []datatypes.JobListing
}
