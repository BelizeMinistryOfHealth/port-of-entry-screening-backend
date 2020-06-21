package domain_test

import (
	"testing"
	"time"

	"bz.epi.covid.screen/arrivals/domain"
)

func TestHydrateCompanions(t *testing.T) {
	dateLayout := "2006-01-02"
	dob, _ := time.Parse(dateLayout, "1977-04-15")
	dateScreened, _ := time.Parse(dateLayout, "2020-05-24")
	arrivalDate, _ := time.Parse(dateLayout, "2020-05-24")
	embarkDate, _ := time.Parse(dateLayout, "2020-05-23")

	newArrival := domain.ArrivalRequest{
		Id:                   "2020-05-28#Kla-Sch#21235345212",
		TravellingCompanions: []string{"2020-05-28#Joh-Bra#212353492"},
		ArrivalInfo: &domain.ArrivalInfo{
			DateOfArrival:        domain.SimpleTime{arrivalDate},
			ModeOfTravel:         "Air",
			VesselNumber:         "A1211",
			CountryOfEmbarkation: "Mexico",
			DateOfEmbarkation:    domain.SimpleTime{embarkDate},
			PortOfEntry:          "pgia",
		},
		PersonalInfo: domain.PersonalInfo{
			FirstName:         "Klara",
			MiddleNameInitial: "U",
			LastName:          "Schuman",
			PassportNumber:    "21235345212",
			Dob:               domain.SimpleTime{dob},
			Nationality:       "Belizean",
		},
		AddressInBelize: &domain.AddressInBelize{
			District:     "cayo",
			Address:      "12 Roseapple Street",
			Municipality: "Belmopan",
		},
		Screening: []domain.Screening{
			{
				FluLikeSymptoms: domain.FluLikeSymptoms{
					Fever:            false,
					Headache:         false,
					Cough:            false,
					Malaise:          false,
					SoreThroat:       false,
					BreathShort:      false,
					BreathDifficulty: false,
					Other:            "",
				},
				OtherSymptoms:             "",
				DiagnosedWithCovid19:      false,
				ContactWithHealthFacility: false,
				Comments:                  "",
				Location:                  "PGIA",
				DateScreened:              domain.SimpleTime{dateScreened},
			},
		},
	}

	companion := domain.ArrivalRequest{
		Id:                   "2020-05-28#Joh-Bra#212353492",
		TravellingCompanions: []string{"2020-05-28#Kla-Sch#21235345212"},
		ArrivalInfo: &domain.ArrivalInfo{
			DateOfArrival:        domain.SimpleTime{arrivalDate},
			ModeOfTravel:         "Air",
			VesselNumber:         "A1211",
			CountryOfEmbarkation: "Mexico",
			DateOfEmbarkation:    domain.SimpleTime{embarkDate},
			PortOfEntry:          "pgia",
		},
		PersonalInfo: domain.PersonalInfo{
			FirstName:      "Johannes",
			LastName:       "Brham",
			PassportNumber: "212353492",
			Dob:            domain.SimpleTime{dob},
			Nationality:    "Belizean",
		},
		AddressInBelize: &domain.AddressInBelize{
			District:     "cayo",
			Address:      "12 Roseapple Street",
			Municipality: "Belmopan",
		},
		Screening: []domain.Screening{
			{
				FluLikeSymptoms: domain.FluLikeSymptoms{
					Fever:            false,
					Headache:         false,
					Cough:            false,
					Malaise:          false,
					SoreThroat:       false,
					BreathShort:      false,
					BreathDifficulty: false,
					Other:            "",
				},
				OtherSymptoms:             "",
				DiagnosedWithCovid19:      false,
				ContactWithHealthFacility: false,
				Comments:                  "",
				Location:                  "PGIA",
				DateScreened:              domain.SimpleTime{dateScreened},
			},
		},
	}

	arrivals := domain.HydrateCompanions([]domain.ArrivalRequest{newArrival, companion})

	first := arrivals[0]
	second := arrivals[1]

	if first.TravellingCompanions[0].FirstName != second.PersonalInfo.FirstName {
		t.Errorf("names did not match")
	}
}
