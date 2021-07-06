package integration

import (
	"bytes"
	"bz.moh.epi/poebackend"
	"net/http"
	"net/http/httptest"
	"testing"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	poebackend.RegistrationFn(w, r)
}

func TestRegistration(t *testing.T) {
	body := `{
    "personalInfo": {
        "firstName": "Roberto",
        "middleName": "Uris",
        "lastName": "Guerra",
        "dob": "1977-04-15",
        "nationality": "BZ",
        "gender": "Male",
        "passportNumber": "000000",
        "phoneNumbers": "000000",
        "occupation": "none",
        "email": "none@mail.com",
        "id": "2021-07-05#Rob-Gue#000000",
        "fullName": "Roberto Uris Guerra",
        "portOfEntry": "PGIA"
    },
    "arrivalInfo": {
        "portOfEntry": "PGIA",
        "dateOfEmbarkation": "2021-07-05T06:00:00.000Z",
        "dateOfArrival": "2021-07-05T06:00:00.000Z",
        "vesselNumber": "AA1211",
        "countryOfEmbarkation": "US",
        "travelOrigin": "Chicago",
        "contactPerson": "Frankie Galleglos",
        "contactPersonPhoneNumber": "1212121",
        "modeOfTravel": "air",
        "purposeOfTrip": "Business",
        "id": "2021-07-05#Rob-Gue#000000"
    },
    "address": {
        "id": "2021-07-05#Rob-Gue#000000",
        "accommodationName": "",
        "address": {
            "id": "ac2142c6-4dad-49b8-9fe7-f89b85de7089",
            "address": "8 Orchid Garden Street",
            "community": {
                "id": "ac2142c6-4dad-49b8-9fe7-f89b85de7089",
                "name": "Belmopan",
                "district": "Cayo"
            }
        }
    },
    "companions": [
        {
            "firstName": "Otto",
            "lastName": "Gallegos",
            "dob": "2013-09-27",
            "nationality": "BZ",
            "gender": "Male",
            "passportNumber": "000000",
            "phoneNumbers": "000000",
            "occupation": "none",
            "email": "none@Mail.com",
            "id": "2021-07-05#Ott-Gal#000000",
            "fullName": "Otto  Gallegos"
        }
    ]
}`

	handler := &myHandler{}
	server := httptest.NewServer(handler)
	client := server.Client()
	defer server.Close()

	resp, err := client.Post(server.URL, "application/json", bytes.NewReader([]byte(body)))
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	t.Logf("resp: %v", resp)

}
