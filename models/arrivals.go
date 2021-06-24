package models

import "time"

// AddressInBelize is the address that the person is residing within Belize
type AddressInBelize struct {
	ID                string     `json:"id" firestore:"id"`
	Address           Address    `json:"address" firestore:"address"`
	ControlID         string     `json:"controlId" firestore:"controlId"`
	AccommodationName string     `json:"accommodationName" firestore:"accommodationName"`
	Modified          *time.Time `json:"modified,omitempty" firestore:"modified"`
	CreatedBy         Editor     `json:"createdBy" firestore:"createdBy"`
	ModifiedBy        Editor     `json:"modifiedBy" firestore:"modifiedBy"`
	Created           *time.Time `json:"created" firestore:"created"`
}

// Community represents a community in a district within Belize
type Community struct {
	ID           string `json:"id" firestore:"id"`
	District     string `json:"district" firestore:"district"`
	Municipality string `json:"name" firestore:"name"`
}

// Address relates the street location or name to a community.
type Address struct {
	Address   string    `json:"address" firestore:"address"`
	Community Community `json:"community" firestore:"community"`
}

// ArrivalInfo contains information specific to an arrival, including port of embarkation and
// vessel information
type ArrivalInfo struct {
	ID                       string     `json:"id" firestore:"id"`
	DateOfArrival            time.Time  `json:"dateOfArrival" firestore:"dateOfArrival"`
	ModeOfTravel             string     `json:"modeOfTravel" firestore:"modeOfTravel"`
	VesselNumber             string     `json:"vesselNumber" firestore:"vesselNumber"`
	CountryOfEmbarkation     string     `json:"countryOfEmbarkation" firestore:"countryOfEmbarkation"`
	DateOfEmbarkation        time.Time  `json:"dateOfEmbarkation" firestore:"dateOfEmbarkation"`
	PortOfEntry              string     `json:"portOfEntry" firestore:"portOfEntry"`
	TravelOrigin             string     `json:"travelOrigin" firestore:"travelOrigin"`
	CountriesVisited         string     `json:"countriesVisited" firestore:"countriesVisited"`
	PurposeOfTrip            string     `json:"purposeOfTrip" firestore:"purposeOfTrip"`
	ContactPerson            string     `json:"contactPerson" firestore:"contactPerson"`
	ContactPersonPhoneNumber string     `json:"contactPersonPhoneNumber" firestore:"contactPersonPhoneNumber"`
	Touch                    bool       `json:"touch" firestore:"touch"`
	Modified                 *time.Time `json:"modified,omitempty" firestore:"modified"`
	CreatedBy                Editor     `json:"createdBy" firestore:"createdBy"`
	ModifiedBy               Editor     `json:"modifiedBy" firestore:"modifiedBy"`
	Created                  *time.Time `json:"created" firestore:"created"`
}

/// Screening Structs
////------------------

// FluLikeSymptoms are the symptoms that a passenger is screened for.
type FluLikeSymptoms struct {
	Fever             bool   `json:"fever" firestore:"fever"`
	Headache          bool   `json:"headache" firestore:"headache"`
	Cough             bool   `json:"cough" firestore:"cough"`
	Malaise           bool   `json:"malaise" firestore:"malaise"`
	SoreThroat        bool   `json:"soreThroat" firestore:"soreThroat"`
	BreathShort       bool   `json:"breathShort" firestore:"breathShort"`
	BreathDifficulty  bool   `json:"breathDifficulty" firestore:"breathDifficulty"`
	RunnyNose         bool   `json:"runnyNose" firestore:"runnyNose"`
	Nausea            bool   `json:"nausea" firestore:"nausea"`
	Diarrhea          bool   `json:"diarrhea" firestore:"diarrhea"`
	ShortnessOfBreath bool   `json:"shortnessOfBreath" firestore:"shortnessOfBreath"`
	Chills            bool   `json:"chills" firestore:"chills"`
	Anosmia           bool   `json:"anosmia" firestore:"anosmia"`
	Aguesia           bool   `json:"aguesia" firestore:"aguesia"`
	Bleeding          bool   `json:"bleeding" firestore:"bleeding"`
	JointMusclePain   bool   `json:"jointMusclePain" firestore:"jointMusclePain"`
	EyeFacialPain     bool   `json:"eyeFacialPain" firestore:"eyeFacialPain"`
	GeneralizedRash   bool   `json:"generalizedRash" firestore:"generalizedRash"`
	BlurredVision     bool   `json:"blurredVision" firestore:"blurredVision"`
	AbdominalPain     bool   `json:"abdominalPain" firestore:"abdominalPain"`
	Vomiting          bool   `json:"vomiting" firestore:"vomiting"`
	Other             string `json:"other" firestore:"other"`
}

// Editor indicates who edited or created a record.
type Editor struct {
	Email string `json:"email" firestore:"email"`
	ID    string `json:"id" firestore:"id"`
}

// Vaccine indicates how many shots are required for a vaccine
type Vaccine struct {
	Name          string `json:"name" firestore:"name"`
	NumberOfShots int    `json:"numberOfShots" firestore:"numberOfShots"`
}

// Vaccination indicates what vaccine a person received and the number of shots
type Vaccination struct {
	Name                 string    `json:"name" firestore:"name"`
	NumberOfShots        int       `json:"numberOfShots" firestore:"numberOfShots"`
	DateOfMostRecentShot time.Time `json:"dateOfMostRecentShot" firestore:"dateOfMostRecentShot"`
}

// Screening is the information collected when screening a person.
type Screening struct {
	ID string `json:"id" firestore:"id"`
	//OtherSymptoms             string          `json:"otherSymptoms" firestore:"otherSymptoms"`
	DiagnosedWithCovid19 bool `json:"diagnosedWithCovid19" firestore:"diagnosedWithCovid19"`
	//ContactWithHealthFacility bool            `json:"contactWithHealthFacility" firestore:"contactWithHealthFacility"`
	Comments                 string          `json:"comments" firestore:"comments"`
	Location                 string          `json:"location" firestore:"location"`
	DateScreened             time.Time       `json:"screened" firestore:"screened"`
	Temperature              float32         `json:"temperature" firestore:"temperature"`
	FluLikeSymptoms          FluLikeSymptoms `json:"fluLikeSymptoms" firestore:"fluLikeSymptoms"`
	TookPcrTestInPast72Hours bool            `json:"tookPcrTestInPast72Hours" firestore:"tookPcrTestInPast72Hours"`
	Vaccination              Vaccination     `json:"vaccination" firestore:"vaccination"`
	Modified                 *time.Time      `json:"modified,omitempty" firestore:"modified"`
	CreatedBy                Editor          `json:"createdBy" firestore:"createdBy"`
	ModifiedBy               Editor          `json:"modifiedBy" firestore:"modifiedBy"`
}

// PersonalInfo about a person, mostly demographic in nature.
type PersonalInfo struct {
	ID                    string    `json:"id" firestore:"id"`
	FirstName             string    `json:"firstName" firestore:"firstName"`
	LastName              string    `json:"lastName" firestore:"lastName"`
	MiddleName            string    `json:"middleName,omitempty" firestore:"middleName" `
	FullName              string    `json:"fullName,omitempty" firestore:"fullName" `
	Dob                   string    `json:"dob" firestore:"dob"`
	Nationality           string    `json:"nationality" firestore:"nationality"`
	PassportNumber        string    `json:"passportNumber,omitempty" firestore:"passportNumber"`
	OtherTravelDocument   string    `json:"otherTravelDocument,omitempty" firestore:"otherTravelDocument"`
	OtherTravelDocumentID string    `json:"otherTravelDocumentId,omitempty" firestore:"otherTravelDocumentId"`
	Email                 string    `json:"email,omitempty" firestore:"email"`
	Gender                string    `json:"gender" firestore:"gender"`
	PhoneNumbers          string    `json:"phoneNumbers,omitempty" firestore:"phoneNumbers"`
	BhisNumber            string    `json:"bhisNumber,omitempty" firestore:"bhisNumber"`
	Occupation            string    `json:"occupation" firestore:"occupation"`
	PortOfEntry           string    `json:"portOfEntry" firestore:"portOfEntry"`
	Created               time.Time `json:"created,omitempty" firestore:"created"`
	Modified              time.Time `json:"modified,omitempty" firestore:"modified"`
	CreatedBy             Editor    `json:"createdBy,omitempty" firestore:"createdBy"`
	ModifiedBy            Editor    `json:"modifiedBy,omitempty" firestore:"modifiedBy"`
}

// TravellingCompanion links a person to a companion. These are usually under age persons travelling with an adult.
type TravellingCompanion struct {
	Relationship string       `json:"relationship" firestore:"relationship"`
	TripID       string       `json:"tripId" firestore:"tripId"`
	PersonID     string       `json:"personId" firestore:"personId"`
	PersonalInfo PersonalInfo `json:"personalInfo" firestore:"personalInfo"`
}

// Arrival is the information related to a specific arrival to the country.
type Arrival struct {
	ID                       string                `json:"id" firestore:"id"`
	ArrivalInfo              ArrivalInfo           `json:"arrivalInfo" firestore:"arrivalInfo"`
	HotelAddress             AddressInBelize       `json:"hotelAddress" firestore:"hotelAddress"`
	Address                  AddressInBelize       `json:"address" firestore:"address"`
	Screenings               []Screening           `json:"screenings" firestore:"screenings"`
	TravellingCompanions     []TravellingCompanion `json:"travellingCompanions" firestore:"travellingCompanions"`
	ContactPerson            string                `json:"contactPerson,omitempty" firestore:"contactPerson"`
	ContactPersonPhoneNumber string                `json:"contactPersonPhoneNumber,omitempty" firestore:"contactPersonPhoneNumber"`
	PurposeOfTrip            string                `json:"purposeOfTrip" firestore:"purposeOfTrip"`
	LengthStay               string                `json:"lengthStay" firestore:"lengthStay"`
	Created                  time.Time             `json:"created" firestore:"created"`
	Modified                 *time.Time            `json:"modified,omitempty" firestore:"modified"`
}

// Person contains personal information and arrival information for a person.
type Person struct {
	ID                 string       `json:"id" firestore:"id"`
	ObjectID           string       `json:"objectID" firestore:"objectID"`
	PersonalInfo       PersonalInfo `json:"personalInfo" firestore:"personalInfo"`
	Arrival            Arrival      `json:"arrivals" firestore:"arrivals"`
	Covid19Vaccination Vaccination  `json:"covid19Vaccination" firestore:"covid19Vaccination"`
	Created            time.Time    `json:"created" firestore:"created"`
	Modified           *time.Time   `json:"modified,omitempty" firestore:"modified"`
	PortOfEntry        string       `json:"portOfEntry" firestore:"portOfEntry"`
}

// WasScreenedOnDate indicates if a person was screened on a specific date.
func (p *Person) WasScreenedOnDate(date time.Time) bool {
	screenings := p.Arrival.Screenings
	for _, s := range screenings {
		if s.DateScreened == date {
			return true
		}
	}
	return false
}
