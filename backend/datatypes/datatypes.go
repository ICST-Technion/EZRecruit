package datatypes

// JobListing represents a job-listing entry.
type JobListing struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Location       string   `json:"location"`
	RequiredSkills []string `json:"requiredSkills"`
}
