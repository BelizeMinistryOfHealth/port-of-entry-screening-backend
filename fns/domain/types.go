package domain

import (
	"fmt"
	"strings"
	"time"
)

type ArrivalSyncStatus int

const (
	NEW_ARRIVAL ArrivalSyncStatus = iota
	SYNCED_ARRIVAL
)

type ScreeningSyncStatus int

const (
	NEW_SCREENING ScreeningSyncStatus = iota
	SYNCED_SCREENING
)

type SimpleTime struct {
	time.Time
}

func (t *SimpleTime) UnmarshalJSON(buf []byte) error {
	layout := "2006-01-02"
	tt, err := time.Parse(layout, strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t *SimpleTime) MarshalJSON() ([]byte, error) {
	layout := "2006-01-02"
	return []byte(`"` + t.Time.Format(layout) + `"`), nil
}

type ArrivalInfo struct {
	DateOfArrival        SimpleTime `json:"date_of_arrival"`
	ModeOfTravel         string     `json:"mode_of_travel"`
	VesselNumber         string     `json:"vessel_number, omitempty"`
	CountryOfEmbarkation string     `json:"country_of_embarkation"`
	DateOfEmbarkation    SimpleTime `json:"date_of_embarkation"`
	PortOfEntry          string     `json:"port_of_entry"`
}

type PersonalInfo struct {
	Id                string     `json:"id, omitempty"`
	FirstName         string     `json:"first_name"`
	MiddleNameInitial string     `json:"middle_name_initial, omitempty"`
	LastName          string     `json:"last_name"`
	FullName          string     `json:"full_name"`
	PassportNumber    string     `json:"passport_number"`
	Dob               SimpleTime `json:"dob"`
	Nationality       string     `json:"nationality"`
}

type AddressInBelize struct {
	District     string `json:"district"`
	Municipality string `json:"municipality"`
	Address      string `json:"address"`
}

type FluLikeSymptoms struct {
	Fever            bool   `json:"fever"`
	Headache         bool   `json:"headache"`
	Cough            bool   `json:"cough"`
	Malaise          bool   `json:"malaise"`
	SoreThroat       bool   `json:"sore_throat"`
	BreathShort      bool   `json:"breath_short"`
	BreathDifficulty bool   `json:"breath_difficulty"`
	Other            string `json:"other"`
}

type Screening struct {
	Id                        string          `json:"id"`
	FluLikeSymptoms           FluLikeSymptoms `json:"flu_like_symptoms"`
	OtherSymptoms             string          `json:"other_symptoms"`
	DiagnosedWithCovid19      bool            `json:"diagnosed_with_covid19"`
	ContactWithHealthFacility bool            `json:"contact_with_health_facility"`
	Comments                  string          `json:"comments, omitempty"`
	Location                  string          `json:"location"`
	DateScreened              SimpleTime      `json:"date_screened"`
}

type ScreeningRequest struct {
	Id                        string              `json:"id"`
	FluLikeSymptoms           FluLikeSymptoms     `json:"flu_like_symptoms"`
	OtherSymptoms             string              `json:"other_symptoms"`
	DiagnosedWithCovid19      bool                `json:"diagnosed_with_covid19"`
	ContactWithHealthFacility bool                `json:"contact_with_health_facility"`
	Comments                  string              `json:"comments, omitempty"`
	Location                  string              `json:"location"`
	DateScreened              SimpleTime          `json:"date_screened"`
	SyncStatus                ScreeningSyncStatus `json:"syncStatus"`
}

type Arrival struct {
	Id                   string           `json:"id"`
	ArrivalInfo          *ArrivalInfo     `json:"arrival_info, omitempty"`
	PersonalInfo         PersonalInfo     `json:"personal_info"`
	AddressInBelize      *AddressInBelize `json:"address_in_belize, omitempty"`
	Screening            []Screening      `json:"screening"`
	TravellingCompanions []PersonalInfo   `json:"travelling_companions, omitempty"`
	Modified             SimpleTime       `json:"modified"`
}

type ArrivalRequest struct {
	Id                   string            `json:"id"`
	ArrivalInfo          *ArrivalInfo      `json:"arrival_info, omitempty"`
	PersonalInfo         PersonalInfo      `json:"personal_info"`
	AddressInBelize      *AddressInBelize  `json:"address_in_Belize, omitempty"`
	Screening            []Screening       `json:"screening"`
	TravellingCompanions []string          `json:"travelling_companions, omitempty"`
	Modified             SimpleTime        `json:"modified"`
	SyncStatus           ArrivalSyncStatus `json:"syncStatus"`
}

func Index(vs []string, t ArrivalRequest) int {
	for i, v := range vs {
		if v == t.Id {
			return i
		}
	}
	return -1
}

func Include(vs []string, t ArrivalRequest) bool {
	return Index(vs, t) >= 0
}

func Filter(vs []ArrivalRequest, f func(ArrivalRequest) bool) []ArrivalRequest {
	vsf := make([]ArrivalRequest, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []ArrivalRequest, f func(arrival ArrivalRequest) PersonalInfo) []PersonalInfo {
	vsm := make([]PersonalInfo, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// MapToPersonalInfo retrieves the PersonalInfo from an array of NewArrivals for the
// corresponding Id.
func MapToPersonalInfo(arrival ArrivalRequest, info []ArrivalRequest) []PersonalInfo {

	arrivals := Filter(info, func(arr ArrivalRequest) bool {
		return arrival.Id != arr.Id && Include(arrival.TravellingCompanions, arr)
	})

	return Map(arrivals, func(arrival ArrivalRequest) PersonalInfo {
		arrival.PersonalInfo.Id = arrival.Id
		arrival.PersonalInfo.FullName = fmt.Sprintf("%s %s", arrival.PersonalInfo.FirstName, arrival.PersonalInfo.LastName)
		return arrival.PersonalInfo
	})

}

// HydrateCompanions takes an array of NewArrivals
// and hydrates their travelling companions with the
// matching PersonalInfo. It returns an array of Arrival
// with the enriched TravellingCompanions.
func HydrateCompanions(arrs []ArrivalRequest) []Arrival {
	arrivals := make([]Arrival, 0)
	for _, v := range arrs {
		infos := MapToPersonalInfo(v, arrs)

		arrival := Arrival{
			Id:                   v.Id,
			ArrivalInfo:          v.ArrivalInfo,
			PersonalInfo:         v.PersonalInfo,
			AddressInBelize:      v.AddressInBelize,
			Screening:            v.Screening,
			Modified:             v.Modified,
			TravellingCompanions: infos,
		}
		arrivals = append(arrivals, arrival)
	}
	return arrivals
}
