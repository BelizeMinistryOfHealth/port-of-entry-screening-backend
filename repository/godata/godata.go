package godata

import (
	"bytes"
	"bz.moh.epi/poebackend/models"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const yes = "Yes"

// Options for making GoData requests
type Options struct {
	Username   string
	Password   string
	Url        string
	Token      string
	OutbreakId string
}

// CaseID Verify if a GoData record exists by fetching
type CaseID struct {
	ID string `json:"id"`
}

type goDataAuthResponse struct {
	AccessToken string `json:"access_token"`
}

// GetGodataToken retrieves a JWT token from a GoData Server.
func GetGodataToken(username, password, baseUrl string) (string, error) {
	reqBody, err := json.Marshal(map[string]string{"username": username, "password": password})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error authenticating with GoData")
		return "", err
	}
	req, err := http.Post(fmt.Sprintf("%s/oauth/token", baseUrl), "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error retrieving token from GoData")
		return "", err
	}

	var authResp *goDataAuthResponse
	if err := json.NewDecoder(req.Body).Decode(&authResp); err != nil {
		log.WithFields(log.Fields{"error": err, "response": req}).Error("failed to decode oauth token from godata")
		return "", err
	}
	return authResp.AccessToken, nil
}

func postToGodata(godataCase GoDataCase, opts Options) (*http.Response, error) {

	token := opts.Token
	outbreakId := opts.OutbreakId

	// Prepare post request to create case
	body, err := json.Marshal(godataCase)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "data": godataCase}).Error("failed to marshal go data case")
		return nil, err
	}

	client := &http.Client{}
	newReq, _ := http.NewRequest("POST", fmt.Sprintf("%s/outbreaks/%s/cases", opts.Url, outbreakId), bytes.NewReader(body))
	newReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	newReq.Header.Set("Content-Type", "application/json")
	log.WithFields(log.Fields{
		"body": body,
		"case": godataCase,
	}).Info("Sending new request to GoData")
	defer newReq.Body.Close()
	return client.Do(newReq)
}

// PostCase creates a new case in a GoData server.
func PostCase(o GoDataCase, opts Options) error {
	resp, err := postToGodata(o, opts)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to post new case to godata")
		return fmt.Errorf("failed to post new case to godata: %+v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"respBody": string(respBody),
		}).WithError(err).Error("Error reading response from godata")
		return fmt.Errorf("error reading godata response: %v  error: %v", respBody, err)
	}
	log.WithFields(log.Fields{
		"case":     o,
		"response": respBody,
	}).Info("Got a response from godata")

	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{"responseFromGoData": respBody}).Error("error posting case to godata")
		return fmt.Errorf("error posting case to godata")
	}

	log.WithFields(
		log.Fields{
			"case":     o,
			"response": respBody,
		}).Info("posted new case to godata")
	return nil
}

// PutCase updates a case in GoData.
func PutCase(o GoDataCase, caseID string, opts Options) error {
	token := opts.Token
	body, err := json.Marshal(o)
	if err != nil {
		return fmt.Errorf("failed to marshal godata case: %w", err)
	}

	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", opts.Url, caseID), bytes.NewReader(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("PutCase() to GoData failed: %w", err)
	}
	defer resp.Body.Close()
	log.WithFields(log.Fields{
		"case":   o,
		"caseID": caseID,
	}).Info("updated godata case")
	return nil
}

// GetCaseByVisualId retrieves a case from GoData that matches the visualId.
// An error is returned if the http request fails or if no case is found.
func GetCaseByVisualId(visualId string, opts Options) (CaseID, error) {
	token := opts.Token
	// We need the id, so we should query for it.
	filter := fmt.Sprintf("{\"where\":{\"visualId\": \"%s\"}}", visualId)
	getUrl := fmt.Sprintf("%s?filter=%s", opts.Url, url.QueryEscape(filter))
	getReq, _ := http.NewRequest(http.MethodGet, getUrl, nil)
	getReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	getReq.Header.Set("Content-Type", "application/json")
	getClient := &http.Client{}
	getResp, err := getClient.Do(getReq)
	if err != nil {
		return CaseID{}, fmt.Errorf("could not retrieve case with visualId: %s | %w", visualId, err)
	}
	defer getResp.Body.Close()
	var resps []CaseID
	if err := json.NewDecoder(getResp.Body).Decode(&resps); err != nil {
		log.WithFields(log.Fields{
			"body":   getResp.Body,
			"status": getResp.Status,
		}).WithError(err).Info("godata case raw body")
		if getResp.StatusCode == http.StatusOK {
			return CaseID{}, nil
		}
		return CaseID{}, fmt.Errorf("failed to decode case data: status %d | godataApiUrl: %s : %w", getResp.StatusCode, getUrl, err)
	}
	if len(resps) == 0 {
		return CaseID{}, fmt.Errorf("no record found with visualId: %s | error: %v", visualId, err)
	}
	return resps[0], nil
}

func eventToGoDataAddress(bzAddress models.AddressInBelize, phoneNumber string, date time.Time) GoDataAddress {
	usualPlaceOfResidence := "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_USUAL_PLACE_OF_RESIDENCE"
	//accommodationResidence := "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_ACCOMMODATION_NAME"

	address := GoDataAddress{
		TypeId:       usualPlaceOfResidence,
		Country:      "Belize",
		City:         bzAddress.Address.Community.Municipality,
		AddressLine1: bzAddress.Address.Address,
		AddressLine2: bzAddress.AccommodationName,
		Date:         date.Format("2006-01-02"),
		PhoneNumber:  phoneNumber,
		LocationId:   bzAddress.Address.Community.ID,
	}
	return address
}

func createGoDataQuestionnaire(screening models.Screening, arrivalInfo models.ArrivalInfo, caseType, purposeOfTrip string) GoDataQuestionnaire {
	var fever = "No"
	if screening.FluLikeSymptoms.Fever {
		fever = yes
	}
	var soreThroat = "No"
	if screening.FluLikeSymptoms.SoreThroat {
		soreThroat = yes
	}
	var cough = "No"
	if screening.FluLikeSymptoms.Cough {
		cough = yes
	}
	var nausea = "No"
	if screening.FluLikeSymptoms.Nausea {
		nausea = yes
	}
	var malaise = "No"
	if screening.FluLikeSymptoms.Malaise {
		malaise = yes
	}
	var runnyNose = "No"
	if screening.FluLikeSymptoms.RunnyNose {
		runnyNose = yes
	}
	var vomiting = "No"
	if screening.FluLikeSymptoms.Vomiting {
		vomiting = yes
	}
	var diarrhea = "No"
	if screening.FluLikeSymptoms.Diarrhea {
		diarrhea = yes
	}
	var shortnessOfBreath = "No"
	if screening.FluLikeSymptoms.ShortnessOfBreath {
		shortnessOfBreath = yes
	}
	var difficultyBreathing = "No"
	if screening.FluLikeSymptoms.BreathDifficulty {
		difficultyBreathing = yes
	}
	var chills = "No"
	if screening.FluLikeSymptoms.Chills {
		chills = yes
	}
	var headache = "No"
	if screening.FluLikeSymptoms.Headache {
		headache = yes
	}
	var anosmia = "No"
	if screening.FluLikeSymptoms.Anosmia {
		anosmia = yes
	}
	var aguesia = "No"
	if screening.FluLikeSymptoms.Aguesia {
		aguesia = yes
	}
	var bleeding = "No"
	if screening.FluLikeSymptoms.Bleeding {
		bleeding = yes
	}
	var jointMusclePain = "No"
	if screening.FluLikeSymptoms.JointMusclePain {
		jointMusclePain = yes
	}
	var eyeFacialPain = "No"
	if screening.FluLikeSymptoms.EyeFacialPain {
		eyeFacialPain = yes
	}
	var generalizedRash = "No"
	if screening.FluLikeSymptoms.GeneralizedRash {
		generalizedRash = yes
	}
	var blurredVision = "No"
	if screening.FluLikeSymptoms.BlurredVision {
		blurredVision = yes
	}
	var abdominalPain = "No"
	if screening.FluLikeSymptoms.AbdominalPain {
		abdominalPain = yes
	}
	var typeOfTraveller = "Non-Tourist"
	if purposeOfTrip == "Vacation" {
		typeOfTraveller = "Tourist"
	}

	var pcrTest = "No"
	if screening.TookPcrTestInPast72Hours {
		pcrTest = yes
	}
	godataQuestionnaire := GoDataQuestionnaire{
		CaseForm: []GoDataCaseForm{
			{Value: []string{
				"Form A0: Minimum data reporting form – for suspected and probable cases",
				"Form A2: Case follow-up form – for confirmed cases (Day 14-21)"},
			},
		},
		DataCollectorName: []DataCollectorName{{Value: screening.CreatedBy.Email}},
		CountryResidence:  nil,
		ShowsSymptoms:     []ShowsSymptoms{{Value: "Unknown"}},
		Fever: []SymptomFever{
			{Value: fever},
		},
		SoreThroat: []SoreThroat{
			{Value: soreThroat},
		},
		RunnyNose: []RunnyNose{
			{Value: runnyNose},
		},
		Cough: []Cough{
			{
				Value: cough,
			},
		},
		Vomiting: []Vomiting{
			{Value: vomiting},
		},
		Nausea: []Nausea{
			{Value: nausea},
		},
		Malaise: []Malaise{
			{Value: malaise},
		},
		Diarrhea: []Diarrhea{
			{Value: diarrhea},
		},
		ShortnessOfBreath: []ShortnessOfBreath{
			{Value: shortnessOfBreath},
		},
		DifficultyBreathing: []DifficultyBreathing{
			{Value: difficultyBreathing},
		},
		SymptomsChills: []SymptomsChills{
			{Value: chills},
		},
		Headache: []Headache{
			{Value: headache},
		},
		Anosmia: []Anosmia{
			{Value: anosmia},
		},
		Aguesia: []Aguesia{
			{Value: aguesia},
		},
		Bleeding: []Bleeding{
			{Value: bleeding},
		},
		JointMusclePain: []JointMusclePain{
			{Value: jointMusclePain},
		},
		EyeFacialPain: []EyeFacialPain{
			{Value: eyeFacialPain},
		},
		GeneralizedRash: []GeneralizedRash{
			{Value: generalizedRash},
		},
		BlurredVision: []BlurredVision{
			{Value: blurredVision},
		},
		AbdominalPain: []AbdominalPain{
			{Value: abdominalPain},
		},
		CaseType: caseType,
		PriorXdayExposureTravelledInternationally: []QuestionnaireAnswer{
			{Value: yes},
		},
		PriorXdayExposureContactWithCase:             nil,
		PriorXDayexposureContactWithCaseDate:         nil,
		PriorXdayExposureInternationalDateTravelFrom: nil,
		PriorXdayExposureInternationalDatetravelTo:   nil,
		PriorXdayexposureInternationaltravelcountries: []PriorXdayExposureInternationalTravelCountries{
			{Value: arrivalInfo.CountriesVisited},
		},
		TypeOfTraveller: []QuestionnaireAnswer{
			{Value: typeOfTraveller},
		},
		PurposeOfTravel: []QuestionnaireAnswer{
			{Value: purposeOfTrip},
		},
		FlightNumber: []QuestionnaireAnswer{
			{Value: arrivalInfo.VesselNumber},
		},
		PriorXdayExposureInternationalTravelCities: nil,
		PcrTestInPast72Hours: []QuestionnaireAnswer{
			{Value: pcrTest},
		},
		PortOfEntry: []QuestionnaireAnswer{{Value: arrivalInfo.PortOfEntry}},
	}

	return godataQuestionnaire
}

func createGoDataCase(goDataQuestionnaire GoDataQuestionnaire, addresses []GoDataAddress, personalInfo models.PersonalInfo, visualId string) GoDataCase {
	return GoDataCase{
		FirstName:                    personalInfo.FirstName,
		MiddleName:                   personalInfo.MiddleName,
		LastName:                     personalInfo.LastName,
		Gender:                       personalInfo.Gender,
		Classification:               "LNG_REFERENCE_DATA_CATEGORY_CASE_CLASSIFICATION_SUSPECT",
		Dob:                          personalInfo.Dob,
		Occupation:                   personalInfo.Occupation,
		DateOfReporting:              time.Now().Format("2006-01-02"),
		IsDateOfReportingApproximate: false,
		Addresses:                    addresses,
		Questionnaire:                goDataQuestionnaire,
		VisualId:                     visualId,
	}
}

type GodataCaseArg struct {
	PersonalInfo models.PersonalInfo
	Screening    models.Screening
	ArrivalInfo  models.ArrivalInfo
	Address      models.AddressInBelize
	VisualId     string
}

func isTourist(purposeOfTrip string) bool {
	nonTouristTypes := []string{"Local", "Resident", "National"}

	for _, a := range nonTouristTypes {
		if a == strings.ToLower(purposeOfTrip) {
			return true
		}
	}
	return false
}

func UpdateGoDataCase(args GodataCaseArg, caseId string, opts Options) error {
	screening := args.Screening
	personalInfo := args.PersonalInfo
	arrivalInfo := args.ArrivalInfo
	visualID := args.VisualId
	address := eventToGoDataAddress(args.Address, personalInfo.PhoneNumbers, arrivalInfo.DateOfArrival)
	var caseType = "Non-Tourist"
	if isTourist(arrivalInfo.PurposeOfTrip) {
		caseType = "Tourist"
	}

	goDataQuestionnaire := createGoDataQuestionnaire(screening, arrivalInfo, caseType, arrivalInfo.PurposeOfTrip)
	godataCase := createGoDataCase(goDataQuestionnaire, []GoDataAddress{address}, personalInfo, visualID)
	log.WithFields(log.Fields{
		"godataCase": godataCase,
		"addresses":  address,
	}).Info("putting to godata")

	return PutCase(godataCase, caseId, opts)
}

func PushToGoData(args GodataCaseArg, opts Options) error {
	screening := args.Screening
	personalInfo := args.PersonalInfo
	arrivalInfo := args.ArrivalInfo
	visualId := args.VisualId
	address := eventToGoDataAddress(args.Address, personalInfo.PhoneNumbers, arrivalInfo.DateOfArrival)
	var caseType = "Non-Tourist"
	if isTourist(arrivalInfo.PurposeOfTrip) {
		caseType = "Tourist"
	}
	goDataQuestionnaire := createGoDataQuestionnaire(screening, arrivalInfo, caseType, arrivalInfo.PurposeOfTrip)
	godataCase := createGoDataCase(goDataQuestionnaire, []GoDataAddress{address}, personalInfo, visualId)
	log.WithFields(log.Fields{
		"godataCase": godataCase,
		"addresses":  address,
	}).Info("Posting to godata")

	return PostCase(godataCase, opts) // push visitor log struct to
}
