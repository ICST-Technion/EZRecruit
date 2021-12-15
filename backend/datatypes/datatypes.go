package datatypes

// JobListing represents a job-listing entry.
type JobListing struct {
	ID             string   `json:"_id"` //wix items need to have an _id field
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Location       string   `json:"location"`
	RequiredSkills []string `json:"requiredSkills"`
	Labels         []string `json:"labels"`
}
