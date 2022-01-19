package datatypes

const (
	statusesCount = 7
)

// TODO: consider exporting through api server, otherwise frontend must be synced manually.
const (
	New                    = iota // חדש
	OfferSent                     // מייל שכר
	PhoneInterview                // ראיון טלפוני
	FirstFrontalInterview         // ראיון פרונטלי ראשון
	SecondFrontalInterview        // ראיון פרונטלי שני
	Task                          // משימה
	Rejected                      // נדחה
	ProposeAnotherJob             // להציע תפקיד אחר
)

// GetLegalStatusIDRange returns bounds on status ID.
func GetLegalStatusIDRange() (int, int) {
	return 0, statusesCount
}

// StatusIDToHebrewString returns status string identifier from int ID.
func StatusIDToHebrewString(status int) string {
	switch status {
	case 0:
		return "חדש"
	case 1:
		return "מייל שכר"
	case 2:
		return "ראיון טלפוני"
	case 3:
		return "ראיון פרונטלי ראשון"
	case 4:
		return "ראיון פרונטלי שני"
	case 5:
		return "משימה"
	case 6:
		return "נדחה"
	case 7:
		return "להציע תפקיד אחר"
	default:
		return "illegal status"
	}
}
