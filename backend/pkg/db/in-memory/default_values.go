package inmemory

import (
	"github.com/ICST-Technion/EZRecruit.git/datatypes"
	"github.com/ICST-Technion/EZRecruit.git/queries"
)

var defaultJobListings = []datatypes.JobListing{
	{
		ID: "0",
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
		ID: "1",
		PostJobListing: queries.PostJobListing{
			Title:       "QA",
			Description: "This is a QA job.",
			Location:    "Tel-Aviv",
			RequiredSkills: []string{
				"TAU Student",
			},
			Labels: []string{
				"QA", "Hi-tech", "Tel-Aviv", "TAU",
			},
		},
	},
	{
		ID: "2",
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
			JobId:     "1",
			User:      "jSmith",
			FirstName: "John",
			LastName:  "Smith",
			Email:     "john@smith-and-sons.com",
			Phone:     "0543333334",
			Labels: []string{
				"job:1", "Generic Name Carrier", "John The Smither", "Programmer", "Technion",
			},
		},
	},
	{
		ID: "1",
		PostJobApplication: queries.PostJobApplication{
			JobId:     "1",
			User:      "Brandon",
			FirstName: "Sleepy",
			LastName:  "Joe",
			Email:     "whitehouse5@email.com",
			Phone:     "0019732956231",
			Labels: []string{
				"job:1", "USA", "White House", "Washington", "Harvard",
			},
		},
	},
	{
		ID: "2",
		PostJobApplication: queries.PostJobApplication{
			JobId:     "1",
			User:      "YantiP",
			FirstName: "Yanti",
			LastName:  "Parazi",
			Email:     "YantiP123@union_of_what_the_hell.com",
			Phone:     "0545533334",
			Labels: []string{
				"job:1", "University of Life", "Dogs", "Sheperd",
			},
		},
	},
	{
		ID: "3",
		PostJobApplication: queries.PostJobApplication{
			JobId:     "0",
			User:      "Delta",
			FirstName: "Corona",
			LastName:  "Virus",
			Email:     "covid19@smith-and-sons.com",
			Phone:     "0541111111",
			Labels: []string{
				"job:0", "China", "Virus", "Bat", "MIT", "Hard working",
			},
		},
	},
	{
		ID: "4",
		PostJobApplication: queries.PostJobApplication{
			JobId:     "2",
			User:      "make_a_trade",
			FirstName: "Ben",
			LastName:  "Simons",
			Email:     "no_phily@yahoo.com",
			Phone:     "0019731632531",
			Labels: []string{
				"job:2", "NBA", "Trade", "Lazy", "Technion",
			},
		},
	},
	{
		ID: "5",
		PostJobApplication: queries.PostJobApplication{
			JobId:     "2",
			User:      "BBB",
			FirstName: "Bomboleo",
			LastName:  "Bombalea",
			Email:     "BBB18329@gmail.com",
			Phone:     "0542733534",
			Labels: []string{
				"job:2", "MIT", "Programmer", "Hard working",
			},
		},
	},
}
