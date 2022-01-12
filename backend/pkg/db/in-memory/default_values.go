package inmemory

import (
	"fmt"

	"github.com/ICST-Technion/EZRecruit/datatypes"
	"github.com/ICST-Technion/EZRecruit/queries"
)

const gForm = "https://docs.google.com/forms/d/e/1FAIpQLSdotIRtdK9qfTQ41ZfmckWV5GiA_c0_TRuIXtMiyJYANLwE3A/viewform"

func getDefaultJobListings() []datatypes.JobListing {
	return []datatypes.JobListing{
		{
			ID: "0",
			PostJobListing: queries.PostJobListing{
				Title:       "Engineer",
				Description: "This is an engineer job.",
				Location:    "Haifa",
				FormLink:    gForm,
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
				FormLink:    gForm,
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
				FormLink:    gForm,
				RequiredSkills: []string{
					"Frontend Dev",
				},
				Labels: []string{
					"Designer", "Designing", "Haifa", "Frontend",
				},
			},
		},
	}
}

func getDefaultApplications() []datatypes.JobApplication {
	return append(getDefaultApplicationsPart1(), getDefaultApplicationsPart2()...)
}

func getStatusLabel(hebrewStatus string) string {
	statusMap := map[string]int{
		"חדש":                 datatypes.New,
		"מייל שכר":            datatypes.OfferSent,
		"ראיון טלפוני":        datatypes.PhoneInterview,
		"ראיון פרונטלי ראשון": datatypes.FirstFrontalInterview,
		"ראיון פרונטלי שני":   datatypes.SecondFrontalInterview,
		"משימה":               datatypes.Task,
		"נדחה":                datatypes.Rejected,
		"להציע תפקיד אחר":     datatypes.ProposeAnotherJob,
	}

	return fmt.Sprintf("status:%d", statusMap[hebrewStatus])
}

// for linters.
func getDefaultApplicationsPart1() []datatypes.JobApplication {
	return []datatypes.JobApplication{
		{
			ID: "0",
			PostJobApplication: queries.PostJobApplication{
				JobId:     "1",
				User:      "jSmith",
				Status:    "חדש",
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
				Status:    "מייל שכר",
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
				Status:    "ראיון טלפוני",
				FirstName: "Yanti",
				LastName:  "Parazi",
				Email:     "YantiP123@union_of_what_the_hell.com",
				Phone:     "0545533334",
				Labels: []string{
					"job:1", "University of Life", "Dogs", "Sheperd",
				},
			},
		},
	}
}

func getDefaultApplicationsPart2() []datatypes.JobApplication {
	return []datatypes.JobApplication{
		{
			ID: "3",
			PostJobApplication: queries.PostJobApplication{
				JobId:     "0",
				User:      "Delta",
				Status:    "ראיון פרונטלי ראשון",
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
				Status:    "ראיון פרונטלי שני",
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
				Status:    "חדש",
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
}
