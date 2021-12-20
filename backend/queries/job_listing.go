package queries

// GetJobListing presents the members of the GetJobListing URL query.
type GetJobListing struct {
	Labels []string `form:"labels" json:"labels"`
}

// PostJobListing presents the members of the POST job-listing request.
type PostJobListing struct {
	Title          string   `form:"title" json:"title"`
	Description    string   `form:"description" json:"description"`
	Location       string   `form:"location" json:"location"`
	RequiredSkills []string `form:"requiredSkills" json:"requiredSkills"`
	Labels         []string `form:"labels" json:"labels"`
}
