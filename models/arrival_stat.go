package models

// ArrivalStat is a document in firesearch that represents some statistics
// about arrivals.
type ArrivalStat struct {
	ID                   string `json:"id"`
	Date                 string `json:"date"`
	Month                string `json:"month"`
	Year                 int    `json:"year"`
	PortOfEntry          string `json:"portOfEntry"`
	CountryOfEmbarkation string `json:"countryOfEmbarkation"`
	PurposeOfTrip        string `json:"purposeOfTrip"`
}
