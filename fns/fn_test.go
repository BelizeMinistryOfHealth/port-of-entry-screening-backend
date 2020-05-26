package fns

import (
	"encoding/json"
	"testing"

	"bz.epi.covid.screen/arrivals/domain"
)

func TestIngest_MarshalPayload(t *testing.T) {
	body := `{
    "id": "1435123123",
	"arrival_info": {
		"date_of_arrival": "2020-05-24",
		"mode_of_travel": "airline",
		"vessel_number": "A3111",
		"country_of_embarkation": "USA",
		"date_of_embarkation": "2020-05-24",
		"prot_of_entry": "pgia"
	},
	"personal_info": {
		"first_name": "Geralt",
		"middle_name_initial": "W",
		"last_name": "Rivia",
		"passport_number": "0012311",
		"dob": "1981-01-01",
		"nationality": "Belizean"
		
	},
	"address_in_belize": {
		"district": "Belize",
		"municipality": "Belize City",
		"address": "12 Albert Street",
		"travelling_companions": []
	},
	"screening": [{
		"flu_like_symptoms": {
			"fever": true,
			"headache": true,
			"cough": false,
			"malaise": false,
			"sore_throat": false,
			"breathShort": false,
			"breath_difficulty": false,
			"other": ""
			},
		"other_symptoms": "",
		"diagnosed_with_covid19": false,
		"contact_with_health_facility": false,
		"comments": "",
		"location": "PGIA",
		"date_screened": "2020-05-24"
	}]
}
`
	var arrival domain.Arrival

	err := json.Unmarshal([]byte(body), &arrival)

	if err != nil {
		t.Fatalf("unmarshall failed")
	}

	t.Log(arrival)
}
