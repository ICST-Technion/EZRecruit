package datatypes

// TODO: consider exporting through api server, otherwise frontend must be synced manually.
const (
	New                    = iota // חדש
	OfferSent                     // מייל שכר
	PhoneInterview                // ראיון טלפוני
	FirstFrontalInterview         // ראיון פרונטלי ראשון
	SecondFrontalInterview        // ראיון פרונטלי שני
	Task                          // משימה
	Rejected                      // נדחה
	ProposeAnotherJob             //להציע תפקיד אחר
)
