package inmemory

import (
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
	"github.com/ICST-Technion/EZRecruit.git/pkg/db"
	"strconv"
)

var defaultJobListings = []datatypes.JobListing{
	{
		ID:          "1",
		Title:       "Engineer",
		Description: "This is an engineer job.",
		Location:    "Haifa",
		RequiredSkills: []string{
			"Can Code", "Technion Student",
		},
		Labels: []string{
			"Engineer", "Technion", "Haifa", "Hi-tech",
		},
	},
	{
		ID:          "2",
		Title:       "Janitor",
		Description: "This is a janitor job.",
		Location:    "Haifa",
		RequiredSkills: []string{
			"TAU Student",
		},
		Labels: []string{
			"Janitor", "Cleaning", "Haifa", "TAU",
		},
	},
	{
		ID:          "3",
		Title:       "Designer",
		Description: "This is a designer job.",
		Location:    "Haifa",
		RequiredSkills: []string{
			"Frontend Dev",
		},
		Labels: []string{
			"Designer", "Designing", "Haifa", "Frontend",
		},
	},
}

// NewInMemoryDB returns a new instance of InMemoryDB.
func NewInMemoryDB() db.DB {
	jobListingsMap := make(map[string]datatypes.JobListing)
	for _, jobListing := range defaultJobListings {
		jobListingsMap[jobListing.ID] = jobListing
	}

	return &DB{
		jobListingsMap: jobListingsMap,
	}
}

// DB struct implements DB interface with in-memory logic.
type DB struct {
	jobListingsMap map[string]datatypes.JobListing
}

// GetJobs returns the stored job-listings.
func (db *DB) GetJobs() []datatypes.JobListing {
	jobListings := make([]datatypes.JobListing, 0)
	for _, listing := range db.jobListingsMap {
		jobListings = append(jobListings, listing)
	}

	return jobListings
}

// InsertJob adds a job to the stored job-listings.
func (db *DB) InsertJob(jobListing *datatypes.JobListing) string {
	if jobListing.ID == "" {
		// assign unique ID0
		jobListing.ID = strconv.Itoa(len(db.jobListingsMap) + 1)
	}
	// rewrite if ID exists
	db.jobListingsMap[jobListing.ID] = *jobListing

	return jobListing.ID
}

// DeleteJob adds a job to the stored job-listings.
func (db *DB) DeleteJob(jobId string) bool {
	if _, found := db.jobListingsMap[jobId]; found {
		delete(db.jobListingsMap, jobId)
		return true
	}

	return false
}
