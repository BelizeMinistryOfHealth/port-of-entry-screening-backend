package models

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPersonalInfo_Decoding(t *testing.T) {
	body := `{
        "id": "2021-06-04#Rob#Rob",
        "firstName": "Roberto",
        "middleName": "Beto",
        "lastName": "Guerra",
        "dob": "1977-04-15T06:00:00.000Z",
        "gender": "Male",
        "nationality": "BZ",
        "passportNumber": "121231231",
        "phoneNumbers": "12312313",
        "occupation": "Engineer",
        "email": "as@as.com",
        "portOfEntry": "PGIA"
    }`

	var personalInfo PersonalInfo
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&personalInfo); err != nil {
		t.Fatalf("decoding personalInfo failed: %v", err)
	}
}

func TestArrivalInfo_Decoding(t *testing.T) {
	body := `{
        "id": "2021-06-04#Rob#Rob",
        "dateOfArrival": "2021-06-04T06:00:00.000Z",
        "dateOfDeparture": "2021-06-12T06:00:00.000Z",
        "dateOfEmbarkation": "2021-06-04T06:00:00.000Z",
        "countryOfEmbarkation": "US",
        "travelOrigin": "Miami",
        "contactPerson": "Me",
        "contactPersonPhoneNumber": "21312313",
        "vesselNumber": "A123132",
        "modeOfTravel": "air",
        "purposeOfTrip": "Tourist",
        "portOfEntry": "PGIA"
    }`

	var arrivalInfo ArrivalInfo
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&arrivalInfo); err != nil {
		t.Fatalf("decoding arrivalInfo failed: %v", err)
	}
}

func TestAddress_Decoding(t *testing.T) {
	body := `{
        "id": "2021-06-04#Rob#Rob",
        "community": {
            "id": "7a60aa72-ab0c-4bbf-9b0f-c55b99d90644",
            "name": "Altun Ha",
            "district": "Belize"
        }
    }`

	var address AddressInBelize
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&address); err != nil {
		t.Fatalf("decoding address failed: %v", err)
	}
}

type RegistrationRequest struct {
	PersonalInfo PersonalInfo    `json:"personalInfo"`
	ArrivalInfo  ArrivalInfo     `json:"arrivalInfo"`
	Address      AddressInBelize `json:"address"`
}

func TestRegistrationRequest_Decode(t *testing.T) {
	body := `{
    "personalInfo": {
        "id": "2021-06-04#Rob#Rob",
        "firstName": "Roberto",
        "middleName": "Beto",
        "lastName": "Guerra",
        "dob": "1977-04-15T06:00:00.000Z",
        "gender": "Male",
        "nationality": "BZ",
        "passportNumber": "121231231",
        "phoneNumbers": "12312313",
        "occupation": "Engineer",
        "email": "as@as.com",
        "portOfEntry": "PGIA"
    },
    "arrivalInfo": {
        "id": "2021-06-04#Rob#Rob",
        "dateOfArrival": "2021-06-04T06:00:00.000Z",
        "dateOfDeparture": "2021-06-12T06:00:00.000Z",
        "dateOfEmbarkation": "2021-06-04T06:00:00.000Z",
        "countryOfEmbarkation": "US",
        "travelOrigin": "Miami",
        "contactPerson": "Me",
        "contactPersonPhoneNumber": "21312313",
        "vesselNumber": "A123132",
        "modeOfTravel": "air",
        "purposeOfTrip": "Tourist",
        "portOfEntry": "PGIA"
    },
    "address": {
        "id": "2021-06-04#Rob#Rob",
        "community": {
            "id": "7a60aa72-ab0c-4bbf-9b0f-c55b99d90644",
            "name": "Altun Ha",
            "district": "Belize"
        }
    }
}`
	var req RegistrationRequest
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&req); err != nil {
		t.Fatalf("failed to decode request: %v", err)
	}

	t.Logf("req: %v", req)
}
