package inmemory

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ICST-Technion/EZRecruit/datatypes"
	"github.com/ICST-Technion/EZRecruit/pkg/db/helpers"
	"github.com/ICST-Technion/EZRecruit/queries"
)

// NewInMemoryDB returns a new instance of InMemoryDB.
func NewInMemoryDB() *DB {
	jobListingsMap := make(map[string]datatypes.JobListing)

	for _, jobListing := range getDefaultJobListings() {
		jobListingsMap[jobListing.ID] = jobListing
	}

	jobApplicationsMap := make(map[string]datatypes.JobApplication)

	for _, jobApplication := range getDefaultApplications() {
		// make label from status
		jobApplication.Labels = append([]string{getStatusLabel(jobApplication.Status)}, jobApplication.Labels...)
		jobApplicationsMap[jobApplication.User] = jobApplication
	}

	return &DB{
		jobListingsMap:         jobListingsMap,
		jobApplicationsMap:     jobApplicationsMap,
		userToResumeMap:        make(map[string]string),
		availableListingID:     len(jobListingsMap),
		availableApplicationID: len(jobApplicationsMap),
	}
}

// DB struct implements DB interface with in-memory logic.
type DB struct {
	jobListingsMap     map[string]datatypes.JobListing
	jobApplicationsMap map[string]datatypes.JobApplication
	userToResumeMap    map[string]string

	availableListingID     int
	availableApplicationID int
}

/// ##############################################
/// ############ JOB LISTING HANDLERS ############
/// ##############################################

// GetJobs function to return stored jobs. If labels is empty then returns all.
func (db *DB) GetJobs(filterable *queries.Filterable, sortable *queries.Sortable) []datatypes.JobListing {
	jobListings := make([]datatypes.JobListing, 0)

	for _, listing := range db.jobListingsMap {
		// filter by
		if len(filterable.FilterLabels) > 0 && !helpers.SetContainsAll(helpers.CreateSetFromSlice(listing.Labels),
			filterable.FilterLabels) {
			continue // not all labels match
		}

		jobListings = append(jobListings, listing)
	}

	// sort result
	sort.Slice(jobListings, func(job1, job2 int) bool {
		intersectionSlice1 := helpers.GetIntersectionSize(jobListings[job1].Labels, sortable.SortLabels)
		intersectionSlice2 := helpers.GetIntersectionSize(jobListings[job2].Labels, sortable.SortLabels)

		// first has more matches
		if intersectionSlice1 > intersectionSlice2 {
			return true
		}
		// first has fewer matches
		if intersectionSlice1 < intersectionSlice2 {
			return false
		}

		return strings.Compare(jobListings[job1].Title, jobListings[job2].Title) > 0
	})

	return jobListings
}

// InsertJob adds a job to the stored job-listings.
func (db *DB) InsertJob(jobListing *datatypes.JobListing) string {
	if jobListing.ID == "" {
		// assign unique ID
		jobListing.ID = strconv.Itoa(db.availableListingID)
		db.availableListingID++
	}
	// (rewrites if ID exists)
	db.jobListingsMap[jobListing.ID] = *jobListing

	return jobListing.ID
}

// DeleteJob adds a job to the stored job-listings.
func (db *DB) DeleteJob(jobID string) bool {
	if _, found := db.jobListingsMap[jobID]; found {
		delete(db.jobListingsMap, jobID)
		return true
	}

	return false
}

/// ##############################################
/// ########## JOB APPLICATION HANDLERS ##########
/// ##############################################

// GetApplications function to return stored jobs. If labels is empty then returns all.
func (db *DB) GetApplications(filterable *queries.Filterable, sortable *queries.Sortable) []datatypes.JobApplication {
	jobApplications := make([]datatypes.JobApplication, 0)

	for _, application := range db.jobApplicationsMap {
		// filter by
		if len(filterable.FilterLabels) > 0 && !helpers.SetContainsAll(helpers.CreateSetFromSlice(application.Labels),
			filterable.FilterLabels) {
			continue // not all labels match
		}

		// get job title for request
		if jobListing, found := db.jobListingsMap[application.JobId]; found {
			application.JobTitle = jobListing.Title
		} else {
			application.JobTitle = "archived"
		}

		jobApplications = append(jobApplications, application)
	}

	// sort result
	sort.Slice(jobApplications, func(job1, job2 int) bool {
		intersectionSlice1 := helpers.GetIntersectionSize(jobApplications[job1].Labels, sortable.SortLabels)
		intersectionSlice2 := helpers.GetIntersectionSize(jobApplications[job2].Labels, sortable.SortLabels)

		// first has more matches
		if intersectionSlice1 > intersectionSlice2 {
			return true
		}
		// first has fewer matches
		if intersectionSlice1 < intersectionSlice2 {
			return false
		}

		return strings.Compare(jobApplications[job1].FirstName, jobApplications[job2].FirstName) > 0
	})

	return jobApplications
}

// InsertApplication adds a job to the stored job-applications.
func (db *DB) InsertApplication(jobApplication *datatypes.JobApplication, resumeLocation string) string {
	if jobApplication.ID == "" {
		// assign unique ID
		jobApplication.ID = strconv.Itoa(db.availableApplicationID)
		db.availableApplicationID++
		// append status ID to labels
		jobApplication.Labels = append([]string{fmt.Sprintf("status:%s", jobApplication.Status)},
			jobApplication.Labels...)
		// append job ID to labels
		jobApplication.Labels = append([]string{fmt.Sprintf("job:%s", jobApplication.JobId)}, jobApplication.Labels...)
	}
	// (rewrites if ID exists)
	db.jobApplicationsMap[jobApplication.User] = *jobApplication

	db.userToResumeMap[jobApplication.User] = resumeLocation

	return jobApplication.ID
}

// SetApplicantsStatus function to update status for given users.
func (db *DB) SetApplicantsStatus(users []string, status int) {
	if min, max := datatypes.GetLegalStatusIDRange(); status < min || status > max {
		return // illegal status
	}

	for _, user := range users {
		if application, found := db.jobApplicationsMap[user]; found {
			application.Status = strconv.Itoa(status)
			// update status label, it is the first (check InsertApplication)
			application.Labels = append([]string{fmt.Sprintf("status:%s", application.Status)},
				application.Labels[1:]...) // size is at least 2 due to the label insertions in insertion func.

			db.jobApplicationsMap[user] = application
		}
	}
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
