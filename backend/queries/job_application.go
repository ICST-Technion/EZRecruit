package queries

// GetJobApplication presents the members of the GetJobApplication URL query.
type GetJobApplication struct {
	Labels []string `form:"labels" json:"labels"`
}

// PostJobApplication presents the members of POST application request.
type PostJobApplication struct {
	JobId    string   `form:"job" json:"job"`
	User     string   `form:"user" json:"user"`
	Name     string   `form:"name" json:"name"`
	Email    string   `form:"email" json:"email"`
	Phone    string   `form:"phone" json:"phone"`
	Location string   `form:"location" json:"location"`
	Labels   []string `form:"labels" json:"labels"`
}
