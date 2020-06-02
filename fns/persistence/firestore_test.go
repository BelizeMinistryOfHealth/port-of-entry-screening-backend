package persistence

import (
	"context"
	"fmt"
	"testing"
	"time"

	"bz.epi.covid.screen/arrivals/domain"
)

func TestFirestoreClient_UpsertArrival(t *testing.T) {
	client := CreateClient(context.Background(), "epi-belize")
	defer client.Close()

	dateLayout := "2006-01-02"
	dob, _ := time.Parse(dateLayout, "1977-04-15")
	dateScreened, _ := time.Parse(dateLayout, "2020-05-24")
	arrivalDate, _ := time.Parse(dateLayout, "2020-05-24")
	embarkDate, _ := time.Parse(dateLayout, "2020-05-23")

	newArrival := domain.NewArrival{
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

	companion := domain.NewArrival{
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

	arrivals := domain.HydrateCompanions([]domain.NewArrival{newArrival, companion})
	err := client.UpsertArrival("port-of-entry-screening", arrivals)

	if err != nil {
		t.Fatalf("UpsertArrival() failed. want nil got %v", err)
	}
}

func TestFirestoreClient_FindByName(t *testing.T) {
	client := CreateClient(context.Background(), "epi-belize")
	defer client.Close()

	res, err := client.FindByLastName("port-of-entry-screening", "Guerra")
	if err != nil {
		t.Fatalf("FindByLastName() error: %v", err)
	}

	if len(res) == 0 {
		t.Fatalf("FindByLastName() result size: got: %d, want: 1", len(res))
	}

	personalInfo := res[0].PersonalInfo

	if personalInfo.FirstName != "Roberto" && personalInfo.LastName != "Guerra" {
		t.Errorf("FindByLastName() got: %s, want: Roberto Guerra",
			fmt.Sprintf("%s %s", personalInfo.FirstName, personalInfo.LastName))
	}

	dob := personalInfo.Dob.Format("2006-01-02")
	if dob != "1977-04-15" {
		t.Errorf("Could not parse dob, got: %s, want: %s", dob, "1977-04-15")
	}

	t.Logf("Results: %v", res)
}

func TestFirestoreClient_FindByScreeningLocation(t *testing.T) {
	client := CreateClient(context.Background(), "epi-belize")
	defer client.Close()

	res, err := client.FindByPortOfEntry("port-of-entry-screening", "pgia")
	if err != nil {
		t.Fatalf("FindByPortOfEntry error: %v", err)
	}

	if len(res) == 0 {
		t.Fatalf("FindByPortOfEntry got: %d, want: 1", len(res))
	}

	portOfEntry := res[0].ArrivalInfo.PortOfEntry

	if portOfEntry != "pgia" {
		t.Errorf("FindByPorfOfEntry got: %s, want: pgia", portOfEntry)
	}
}

func TestPaginateQuery(t *testing.T) {
	client := CreateClient(context.Background(), "epi-belize")
	defer client.Close()

	arrivals, err := client.List("port-of-entry-screening", "2020-05-23%23Mau-Xhi%2399119911")

	if err != nil {
		t.Fatalf("error listing arrivals: %+v", err)
	}

	t.Logf("Arrivals: %+v", arrivals)
}

func TestFirestoreClient_SaveScreening(t *testing.T) {
	client := CreateClient(context.Background(), "epi-belize")
	defer client.Close()

	dateLayout := "2006-01-02"
	dateScreened, _ := time.Parse(dateLayout, "2020-06-01")
	screening := domain.Screening{
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
		Location:                  "Radisson Hotel",
		DateScreened:              domain.SimpleTime{dateScreened},
	}
	err := client.SaveScreening("port-of-entry-screening", "2020-05-30#Mig-Sag#97119911", screening)

	if err != nil {
		t.Errorf("SaveScreening() failed: %v", err)
	}
}
