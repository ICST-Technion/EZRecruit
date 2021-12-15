package inmemory

import (
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
	"github.com/ICST-Technion/EZRecruit.git/pkg/db"
)

var defaultJobListings = []datatypes.JobListing{
	{
		ID:          "job1",
		Title:       "Engineer",
		Description: "This is an engineer job.",
		Location:    "Haifa",
		RequiredSkills: []string{
			"Can Code",
			"Technion Student",
		}},
}

// NewInMemoryDB returns a new instance of InMemoryDB.
func NewInMemoryDB() db.DB {
	return &DB{
		jobListings: defaultJobListings,
	}
}

// DB struct implements DB interface with in-memory logic.
type DB struct {
	jobListings []datatypes.JobListing
}

// GetJobs returns the stored job-listings.
func (db *DB) GetJobs() []datatypes.JobListing {
	return db.jobListings
}

// AppendJob adds a job to the stored job-listings.
func (db *DB) AppendJob(jobListing datatypes.JobListing) {
	db.jobListings = append(db.jobListings, jobListing)
}
