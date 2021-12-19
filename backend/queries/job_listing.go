package queries

// GetJobListingQuery presents the members of the GetJobListing URL query.
type GetJobListingQuery struct {
	Labels []string `form:"labels"`
}
