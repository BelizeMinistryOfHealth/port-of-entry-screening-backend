package sib

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firesearch"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
)

const isoDateLayout = "2006-01-02"

// PurposeOfTrip is the list of purpose of visit in the Travel App.
//nolint
var PurposeOfTrip = map[string]string{
	"1":  "Vacation",
	"2":  "Sport",
	"3":  "Official",
	"4":  "Business",
	"5":  "Study",
	"6":  "Visiting Friends/Relatives",
	"7":  "Honeymoon/Wedding",
	"8":  "Meeting/Convention",
	"9":  "Intransit",
	"10": "Other",
	"11": "Second Home Owner",
	"12": "Repatriate",
}

// Port is the port of entry that is represented in the Travel App.
// The Travel App uses numbers to represent the ports of entry.
type Port string

const (
	// PGIA Port of Entry
	PGIA Port = "1"
)

// ToPortOfEntry converts an sib port of entry to a port of entry that the POE understands.
func ToPortOfEntry(p Port) string {
	switch p {
	case PGIA:
		return "PGIA"
	default:
		return "PGIA"
	}
}

// TravelMode represents the travel modes in the Travel App.
// The Travel App only supports `air` travel mode.
type TravelMode string

const (
	// Air travel mode
	Air TravelMode = "1"
)

// ToTravelMode converts a TravelMode as is in the Travel App, to a readable string
func ToTravelMode(t TravelMode) string {
	switch t {
	case Air:
		return "Air"
	default:
		return "Land"
	}
}

// District is a numeric representation of a district
type District string

const (
	// Corozal District
	Corozal District = "1"
	// OrangeWalk District
	OrangeWalk District = "2"
	// Belize District
	Belize District = "3"
	// Cayo District
	Cayo District = "4"
	// StannCreek District
	StannCreek District = "5"
	// Toledo District
	Toledo District = "6"
)

// ToDistrict converts a district numeric code to a readable string
func ToDistrict(d District) string {
	switch d {
	case Corozal:
		return "Corozal"
	case OrangeWalk:
		return "Orange Walk"
	case Belize:
		return "Belize"
	case Cayo:
		return "Cayo"
	case StannCreek:
		return "Stann Creek"
	case Toledo:
		return "Toledo"
	default:
		return ""
	}
}

// Gender encodes a person's gender as represented in the Travel App.
type Gender string

const (
	// Male gender
	Male Gender = "1"
	// Female gender
	Female Gender = "2"
)

// TravelStatus is a numeric representation of a person's travel as encoded in the Travel App.
type TravelStatus string

const (
	// Default travel status
	Default TravelStatus = "0"
	// Cleared through screening
	Cleared TravelStatus = "1"
	// NotCleared through screening
	NotCleared TravelStatus = "2"
	// Departed the country
	Departed TravelStatus = "3"
)

// Arrival is a representation of the record that SIB's endpoint
// returns. This does not match our own internal structure, but it
// facilitates unmarshalling so that we can convert it to our own
// Person model.
type Arrival struct {
	ID                  string       `json:"id"`
	TripID              string       `json:"tripId"`
	Status              TravelStatus `json:"status"`
	FirstName           string       `json:"fname"`
	MiddleName          string       `json:"mname"`
	LastName            string       `json:"lname"`
	Email               string       `json:"email"`
	Gender              string       `json:"sex"`
	Nationality         string       `json:"nationality"`
	PassportNumber      string       `json:"passnum"`
	PhoneNumber         string       `json:"phone"`
	ContactPerson       string       `json:"contactPer"`
	ContactPersonNumber string       `json:"contactPerNum"`
	TravelMode          TravelMode   `json:"travelMode"`
	VesselNumber        string       `json:"vesselNum"`
	PortOfEntry         Port         `json:"port"`
	TravelOrigin        string       `json:"travelOrigin"`
	CountryVisited      string       `json:"countryVisited"`
	DateOfEmbarkation   string       `json:"dateEmbarktion"`
	TravelDate          string       `json:"travelDate"`
	CityAirport         string       `json:"cityAirport"`
	Dob                 string       `json:"dateOfBirth"`
	DateCreated         string       `json:"createdAt"`
	District            District     `json:"facilityDistrict"`
	Community           string       `json:"facilityCTV"`
	TouristMainAddress  string       `json:"touristMainAddress"`
	PurposeOfTrip       string       `json:"purposeOfTrip"`
	Occupation          string       `json:"occupation"`
	LengthStay          string       `json:"lengthStay"`
}

func (s *Arrival) generateFullName() string {
	var fullname = fmt.Sprintf("%s %s", s.FirstName, s.LastName)
	if len(s.MiddleName) > 0 {
		fullname = fmt.Sprintf("%s %s %s", s.FirstName, s.MiddleName, s.LastName)
	}
	return fullname
}

// ToAddressInBelize converts the address as represented in the Travel App to the format expected in POE.
func ToAddressInBelize(ctx context.Context, locationsStore *firesearch.LocationStore, s Arrival) (models.AddressInBelize, error) {
	if len(s.District) > 0 {
		district := ToDistrict(s.District)
		communityName := s.Community
		community, err := locationsStore.FindByName(ctx, communityName, district)
		if err != nil {
			log.WithFields(log.Fields{"sibArrival": s}).WithError(err).Error("failed to query for community")
			return models.AddressInBelize{}, fmt.Errorf("ToAddressInBelize error: %w", err)
		}

		address := models.AddressInBelize{
			Address: models.Address{
				Community: models.Community{
					ID:           community.ID,
					District:     district,
					Municipality: community.Name,
				},
			},
			ControlID:         "",
			AccommodationName: "",
		}
		return address, nil
	}
	return models.AddressInBelize{}, nil
}

// HotelsToAddressInBelize converts the hotel represented in the Travel App to the format expected in POE.
func HotelsToAddressInBelize(ctx context.Context, hotelStore firesearch.HotelStore, s Arrival) (models.AddressInBelize, error) {
	a := s.TouristMainAddress
	if strings.Index(a, "[") == 0 {
		var addressWithNoBrackets = strings.Replace(a, "[", "", -1)
		addressWithNoBrackets = strings.Replace(addressWithNoBrackets, "]", "", -1)
		var addressWithNoBraces = strings.Replace(addressWithNoBrackets, "{", "", -1)
		addressWithNoBraces = strings.Replace(addressWithNoBraces, "}", "", -1)
		addressParts := strings.Split(addressWithNoBraces, ",")
		controlID := strings.Split(addressParts[0], "=")
		if len(controlID) != 2 {
			return models.AddressInBelize{}, nil
		}

		address, err := hotelStore.FindHotelByID(ctx, controlID[1])
		if err != nil {
			return models.AddressInBelize{}, fmt.Errorf("HotelsToAddressInBelize failed: %w", err)
		}

		log.WithFields(log.Fields{
			"controlID": controlID,
			"address":   address,
		}).Info("Found address for controlID")

		return address, nil

	}

	if a == "0" {
		return models.AddressInBelize{}, nil
	}

	address, err := hotelStore.FindHotelByID(ctx, a)
	if err != nil {
		return models.AddressInBelize{}, fmt.Errorf("HotelsToAddressInBelize failed: %w", err)
	}

	log.WithFields(log.Fields{
		"controlId": a,
		"address":   address,
	}).Info("Found address for controlId")

	return address, nil
}

// ToGender converts the numeric representation of a gender to a readable string as expected by POE
func ToGender(g string) string {
	// Travel App is not consistent with how the gender data is presented.In some cases it is a code where 1 = male,
	// 2 == female.In other cases it has the format: "[{label=Female, value=2}]"
	if g == "1" {
		return "Male"
	}

	if g == "2" {
		return "Female"
	}
	var genderWithNoBrackets = strings.Replace(g, "[", "", -1)
	genderWithNoBrackets = strings.Replace(genderWithNoBrackets, "]", "", -1)
	var genderWithNoBraces = strings.Replace(genderWithNoBrackets, "{", "", -1)
	genderWithNoBraces = strings.Replace(genderWithNoBraces, "}", "", -1)
	genderParts := strings.Split(genderWithNoBraces, ",")
	genderValue := strings.Split(genderParts[0], "=")
	return genderValue[1]
}

// ToPerson converts an SIB Arrival record to a record expected by POE
func (s *Arrival) ToPerson() models.Person {
	nationality := FindCountryByName(s.Nationality)
	personID := s.GenerateID()
	travelDate, _ := time.Parse(isoDateLayout, s.TravelDate)
	dateCreated, _ := time.Parse(isoDateLayout, s.DateCreated)

	person := models.Person{
		ID: personID,
		PersonalInfo: models.PersonalInfo{
			FirstName:             s.FirstName,
			LastName:              s.LastName,
			MiddleName:            s.MiddleName,
			FullName:              s.generateFullName(),
			Dob:                   s.Dob,
			Nationality:           nationality,
			PassportNumber:        s.PassportNumber,
			OtherTravelDocument:   "",
			OtherTravelDocumentID: "",
			Email:                 s.Email,
			Gender:                ToGender(s.Gender),
			PhoneNumbers:          s.PhoneNumber,
			BhisNumber:            "",
			Occupation:            s.Occupation,
		},
		Arrival: models.Arrival{
			ArrivalInfo: models.ArrivalInfo{
				DateOfArrival:        travelDate,
				ModeOfTravel:         ToTravelMode(s.TravelMode),
				VesselNumber:         s.VesselNumber,
				CountryOfEmbarkation: FindCountryByName(s.TravelOrigin),
				DateOfEmbarkation:    s.DateOfEmbarkation,
				PortOfEntry:          ToPortOfEntry(s.PortOfEntry),
				TravelOrigin:         s.CityAirport,
				CountriesVisited:     s.CountryVisited,
			},
			//Addresses:                addresses,
			Screenings:               nil,
			TravellingCompanions:     nil,
			ContactPerson:            s.ContactPerson,
			ContactPersonPhoneNumber: s.ContactPersonNumber,
			PurposeOfTrip:            PurposeOfTrip[s.PurposeOfTrip],
			LengthStay:               s.LengthStay,
			Created:                  dateCreated,
			Modified:                 &dateCreated,
		},
		Created:     dateCreated,
		Modified:    &dateCreated,
		PortOfEntry: ToPortOfEntry(s.PortOfEntry),
	}
	return person
}

// PayloadItems represents the format of how the items are encoded in the TravellersData
type PayloadItems struct {
	ID        string    `json:"id"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Arrivals  struct {
		Items []Arrival `json:"items"`
	} `json:"arrivals"`
}

// TravellersData represents the format of data in the Travel App
type TravellersData struct {
	Data struct {
		ListPersons struct {
			Items     []PayloadItems `json:"items"`
			NextToken string         `json:"nextToken"`
		} `json:"listPersonsByDate"`
	} `json:"data"`
}

// GenerateID generates an ID for the Person record
// that will be created. The ID is a composite of
// travelDate + name + passport number
func (s *Arrival) GenerateID() string {
	reg, _ := regexp.Compile(`[^a-zA-Z0-9]+`)

	fname := reg.ReplaceAllString(s.FirstName, " ")
	if len(fname) > 2 {
		fname = fname[:3]
	}
	lname := reg.ReplaceAllString(s.LastName, " ")
	if len(lname) > 2 {
		lname = lname[:3]
	}
	return fmt.Sprintf("%s#%s-%s#%s", s.TravelDate[:10], fname, lname, s.PassportNumber)
}
