# Port of Entry COVID10 Screening App
The Backend of the COVID19 Port Of Entry Screening App

## Data Model
The data model can be subdivided into 4 general areas:
1. Arrival Info
2. Passenger Info
3. Address In Belize
4. Screening
5. Travelling Companions

The travelling companions arrival information is the same. The client is responsible
for gathering the unique information (screening and passenger information). It should
create a separate record for each companion and associate the related travelling companions.

For example, if traveller A is accompanied by B, C and D. Then A's travelling companions
should be [B, C, D]. And B's travelliing companions should be [A, C, D]. And C's travelling companions
should be [A, B, D]; and D's travelling companions should be [A, B, C].

## Infrastructure Overview

This system provides the backend for the port of entry COVID19 screening for Belize.
The APIs are purposefully separate from any front end application so that we can enable
inter-agency data sharing and quick access to the raw data to the EPI Unit.
It also allows us to use relatively cost-efficient methods of deployment. Using the serverless pattern
allows us to scale our servers to 0 and only be billed for the compute time used and not idle time. 
This allows us to iterate and improvise as fast as the situation changes.

## The API

To get access to the API, contact a member of the EPI unit at the Ministry of Health.
Every request to the API `must` have a `client_id` and a `client_secret` in the HTTP Header.

### Recording Arrivals

POST /ingest

Sample body:

```json
[{
	"id": "0012311",
	"travelling_companions": [],
	"arrival_info": {
		"date_of_arrival": "2020-05-24",
		"mode_of_travel": "airline",
		"vessel_number": "A3111",
		"country_of_embarkation": "USA",
		"date_of_embarkation": "2020-05-24",
		"port_of_entry": "pgia"
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
		"address": "12 Albert Street"
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
}]
```
The client `must` generate a unique ID. Posting to the same URL will create a record or update the existing
one if there is a record with the `id` provided.
 

### Retrieving Arrivals by Port of Entry
Arrivals at a port of entry can be retrieved with the use of a query string:

__GET /byPortOfEntry?port_entry=pgia__

The response is an array of arrivals at that port of entry:

```json
[
    {
        "id": "0012311",
        "arrival_info": {
            "date_of_arrival": "2020-05-24",
            "mode_of_travel": "airline",
            "vessel_number": "A3111",
            "country_of_embarkation": "USA",
            "date_of_embarkation": "2020-05-24",
            "port_of_entry": "pgia"
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
            "address": "12 Albert Street"
        },
        "screening": [
            {
                "flu_like_symptoms": {
                    "fever": true,
                    "headache": true,
                    "cough": false,
                    "malaise": false,
                    "sore_throat": false,
                    "breath_short": false,
                    "breath_difficulty": false,
                    "other": ""
                },
                "other_symptoms": "",
                "diagnosed_with_covid19": false,
                "contact_with_health_facility": false,
                "comments": "",
                "location": "PGIA",
                "date_screened": "2020-05-24"
            }
        ],
        "travelling_companions": null
    },
    {
        "id": "",
        "arrival_info": null,
        "personal_info": {
            "first_name": "",
            "middle_name_initial": "",
            "last_name": "",
            "passport_number": "",
            "dob": "0001-01-01",
            "nationality": ""
        },
        "address_in_belize": null,
        "screening": [
            {
                "flu_like_symptoms": {
                    "fever": false,
                    "headache": false,
                    "cough": false,
                    "malaise": false,
                    "sore_throat": false,
                    "breath_short": false,
                    "breath_difficulty": false,
                    "other": ""
                },
                "other_symptoms": "",
                "diagnosed_with_covid19": false,
                "contact_with_health_facility": false,
                "comments": "",
                "location": "",
                "date_screened": "0001-01-01"
            }
        ],
        "travelling_companions": null
    }
]
```
