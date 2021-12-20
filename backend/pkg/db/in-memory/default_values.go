package inmemory

import (
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
	"github.com/ICST-Technion/EZRecruit.git/queries"
)

var defaultJobListings = []datatypes.JobListing{
	{
		ID: "1",
		PostJobListing: queries.PostJobListing{
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
	},
	{
		ID: "2",
		PostJobListing: queries.PostJobListing{
			Title:       "Janitor",
			Description: "This is a janitor job.",
			Location:    "Tel-Aviv",
			RequiredSkills: []string{
				"TAU Student",
			},
			Labels: []string{
				"Janitor", "Cleaning", "Tel-Aviv", "TAU",
			},
		},
	},
	{
		ID: "3",
		PostJobListing: queries.PostJobListing{
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
	},
}

var defaultJobApplications = []datatypes.JobApplication{
	{
		ID: "0",
		PostJobApplication: queries.PostJobApplication{
			JobId:    "1",
			User:     "jSmith",
			Name:     "John Smith",
			Email:    "john@smith-and-sons.com",
			Phone:    "0543333334",
			Location: "Haifa",
			Labels: []string{
				"Generic Name Carrier", "John The Smither", "Programmer", "Technion",
			},
		},
	},
}
