package godata

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_GetCaseByVisualId(t *testing.T) {

	response := `
[
  {
    "firstName": "Xxx",
    "gender": "Male",
    "wasContact": false,
    "safeBurial": false,
    "classification": "LNG_REFERENCE_DATA_CATEGORY_CASE_CLASSIFICATION_SUSPECT",
    "transferRefused": false,
    "questionnaireAnswers": {
      "Case_WhichForm": [
        {
          "value": [
            "Form A0: Minimum data reporting form – for suspected and probable cases",
            "Form A2: Case follow-up form – for confirmed cases (Day 14-21)"
          ]
        }
      ],
      "FA0_datacollector_name": [
        {
          "value": ""
        }
      ],
      "FA0_case_countryresidence": null,
      "FA0_symptoms_caseshowssymptoms": [
        {
          "value": "Unknown"
        }
      ],
      "FA0_symptom_fever": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_sorethroat": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_runnynose": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_cough": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_vomiting": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_nausea": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_diarrhea": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_shortnessofbreath": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_difficulty_breathing": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_chills": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_headache": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_malaise": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_anosmia": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_aguesia": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_bleeding": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_joint_muscle_pain": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_eye_facial_pain": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_generalized_rash": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_blurred_vision": [
        {
          "value": "No"
        }
      ],
      "FA0_symptom_abdominal_pain": [
        {
          "value": "No"
        }
      ],
      "case_type": "Tourist",
      "FA0_priorXdayexposure_travelledinternationally": [
        {
          "value": "Yes"
        }
      ],
      "FA0_priorXdayexposure_contactwithcase": null,
      "FA0_priorXdayexposure_contactwithcasedate": null,
      "FA0_priorXdayexposure_internationaldatetravelfrom": null,
      "FA0_priorXdayexposure_internationaldatetravelto": null,
      "FA0_priorXdayexposure_internationaltravelcountries": [
        {
          "value": ""
        }
      ],
      "FA0_priorXdayexposure_internationaltravelcities": null,
      "FA0_priorXdayexposure_typeoftraveler": [
        {
          "value": "Tourist"
        }
      ],
      "FA0_priorXdayexposure_purposeoftravel": [
        {
          "value": "Tourist"
        }
      ],
      "FA0_priorXdayexposure_flightnumber": [
        {
          "value": "Dl1983"
        }
      ],
      "FA0_priorXdayexposure_tookpcrtest_past72hours": [
        {
          "value": "No"
        }
      ],
      "port_of_entry": [
        {
          "value": "PGIA"
        }
      ]
    },
    "id": "83d63d59-e682-4c3f-8279-e525ff1fe6c4",
    "outbreakId": "XXX-XXX-XXX",
    "visualId": "PGIA#2021-06-10#Xxx-Xxx#XXXXXXX",
    "middleName": "Xxx",
    "lastName": "Xxx",
    "dob": "1962-03-28T00:00:00.000Z",
    "occupation": "Contractor ",
    "addresses": [
      {
        "typeId": "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_USUAL_PLACE_OF_RESIDENCE",
        "country": null,
        "city": "",
        "addressLine1": "Grand Caribe",
        "postalCode": null,
        "locationId": "91dc8871-a878-439b-9352-0ffdcf95f42a",
        "geoLocationAccurate": false,
        "date": "2021-06-10T00:00:00.000Z",
        "phoneNumber": "8657223526",
        "emailAddress": null
      }
    ],
    "dateOfReporting": "2021-06-10T00:00:00.000Z",
    "isDateOfReportingApproximate": false,
    "classificationHistory": [
      {
        "classification": "LNG_REFERENCE_DATA_CATEGORY_CASE_CLASSIFICATION_SUSPECT",
        "startDate": "2021-06-10T18:20:32.545Z",
        "endDate": null
      }
    ],
    "hasRelationships": false,
    "usualPlaceOfResidenceLocationId": "91dc8871-a878-439b-9352-0ffdcf95f42a",
    "createdAt": "2021-06-10T18:20:32.546Z",
    "createdBy": "7776e3af-1872-43f8-9cfc-f8add6f617b7",
    "updatedAt": "2021-06-11T00:35:26.329Z",
    "updatedBy": "7776e3af-1872-43f8-9cfc-f8add6f617b7",
    "createdOn": "API",
    "deleted": false,
    "address": {}
  }
]`

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(response))
		if err != nil {
			t.Fatalf("error in test server: %v", err)
		}
	}))

	defer server.Close()

	visualID := "PGIA#2021-06-10#Xxx-Xxx#Xxxxx"
	opts := Options{
		Username:   "Xxx.xxx@xxx.net",
		Password:   "Xxxxx",
		URL:        server.URL,
		Token:      "zDU31XGBhGxXfgdz4TUc0xgKBwjKo5otBwG79Db5D7GUqWrwBH1V49K3qyvOJggZ",
		OutbreakID: "5fc2d66b-8af8-42eb-a47a-c56fdd42264a",
	}

	api := NewApi(opts.URL, server.Client())

	caseID, err := api.GetCaseByVisualId(visualID, opts)
	require.NoError(t, err)

	assert.Equal(t, "83d63d59-e682-4c3f-8279-e525ff1fe6c4", caseID.ID)
}

func TestApi_GetCaseByVisualId_NoResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("[]"))
	}))

	defer server.Close()
	visualID := "PGIA#2021-06-10#Xxx-Xxx#Xxxxx"
	opts := Options{
		Username:   "Xxx.xxx@xxx.net",
		Password:   "Xxxxx",
		URL:        server.URL,
		Token:      "zDU31XGBhGxXfgdz4TUc0xgKBwjKo5otBwG79Db5D7GUqWrwBH1V49K3qyvOJggZ",
		OutbreakID: "5fc2d66b-8af8-42eb-a47a-c56fdd42264a",
	}
	api := NewApi(opts.URL, server.Client())

	caseID, err := api.GetCaseByVisualId(visualID, opts)
	var noResultsErr *NoResultsErr
	assert.ErrorAs(t, err, &noResultsErr)
	assert.Equal(t, "", caseID.ID)

}
