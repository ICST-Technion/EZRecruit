package inmemory

import (
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
	"github.com/ICST-Technion/EZRecruit.git/pkg/db"
	"github.com/ICST-Technion/EZRecruit.git/pkg/db/helpers"
	"strconv"
)

// NewInMemoryDB returns a new instance of InMemoryDB.
func NewInMemoryDB() db.DB {
	jobListingsMap := make(map[string]datatypes.JobListing)
	for _, jobListing := range defaultJobListings {
		jobListingsMap[jobListing.ID] = jobListing
	}

	jobApplicationsMap := make(map[string]datatypes.JobApplication)
	for _, jobApplication := range defaultJobApplications {
		jobApplicationsMap[jobApplication.User] = jobApplication
	}

	return &DB{
		jobListingsMap:         jobListingsMap,
		jobApplicationsMap:     jobApplicationsMap,
		userToResumeMap:        make(map[string]string),
		availableListingId:     len(jobListingsMap),
		availableApplicationId: len(jobApplicationsMap),
	}
}

// DB struct implements DB interface with in-memory logic.
type DB struct {
	jobListingsMap     map[string]datatypes.JobListing
	jobApplicationsMap map[string]datatypes.JobApplication
	userToResumeMap    map[string]string

	availableListingId     int
	availableApplicationId int
}

/// ##############################################
/// ############ JOB LISTING HANDLERS ############
/// ##############################################

// GetJobs function to return stored jobs. If labels is empty then returns all.
func (db *DB) GetJobs(labels []string) []datatypes.JobListing {
	jobListings := make([]datatypes.JobListing, 0)
	for _, listing := range db.jobListingsMap {
		if len(labels) > 0 && !helpers.SetContainsAll(helpers.CreateSetFromSlice(listing.Labels), labels) {
			continue // not all labels match
		}

		jobListings = append(jobListings, listing)
	}

	return jobListings
}

// InsertJob adds a job to the stored job-listings.
func (db *DB) InsertJob(jobListing *datatypes.JobListing) string {
	if jobListing.ID == "" {
		// assign unique ID
		jobListing.ID = strconv.Itoa(db.availableListingId)
		db.availableListingId++
	}
	// (rewrites if ID exists)
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

/// ##############################################
/// ########## JOB APPLICATION HANDLERS ##########
/// ##############################################

// GetApplications function to return stored jobs. If labels is empty then returns all.
func (db *DB) GetApplications(labels []string) []datatypes.JobApplication {
	jobApplications := make([]datatypes.JobApplication, 0)
	for _, application := range db.jobApplicationsMap {
		if len(labels) > 0 && !helpers.SetContainsAll(helpers.CreateSetFromSlice(application.Labels), labels) {
			continue // not all labels match
		}

		// get job title for request
		if jobListing, found := db.jobListingsMap[application.JobId]; found {
			application.JobTitle = jobListing.Title
			jobApplications = append(jobApplications, application)
		}
	}

	return jobApplications
}

// InsertApplication adds a job to the stored job-applications.
func (db *DB) InsertApplication(jobApplication *datatypes.JobApplication, resumeLocation string) string {
	if jobApplication.ID == "" {
		// assign unique ID
		jobApplication.ID = strconv.Itoa(db.availableApplicationId)
		db.availableApplicationId++
	}
	// (rewrites if ID exists)
	db.jobApplicationsMap[jobApplication.User] = *jobApplication

	db.userToResumeMap[jobApplication.User] = resumeLocation

	return jobApplication.ID
}

/// ##############################################
/// ############ RESUME FILE HANDLERS ############
/// ##############################################

// GetUserResumeFileLocation returns user's resume file location.
func (db *DB) GetUserResumeFileLocation(user string) (string, bool) {
	location, found := db.userToResumeMap[user]
	if !found {
		location = "default application"
	}

	return location, found
}
