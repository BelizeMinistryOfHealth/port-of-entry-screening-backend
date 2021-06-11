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

// CaseArg are the arguments for creating a new GoData Case
type CaseArg struct {
	PersonalInfo models.PersonalInfo
	Screening    models.Screening
	ArrivalInfo  models.ArrivalInfo
	Address      models.AddressInBelize
	VisualID     string
}

// Options for making GoData requests
type Options struct {
	Username   string
	Password   string
	URL        string
	Token      string
	OutbreakID string
}

// CaseID Verify if a GoData record exists by fetching
type CaseID struct {
	ID string `json:"id"`
}

type goDataAuthResponse struct {
	AccessToken string `json:"access_token"`
}

// API is the GoData api
type API interface {
	GetCaseByVisualId(visualID string, opts Options) (CaseID, error)
	UpdateCase(args CaseArg, caseID string, opts Options) error
	CreateCase(args CaseArg, opts Options) error
}

type api struct {
	Client  *http.Client
	baseURL string
}

// NewAPI creates a new GoData API
func NewAPI(baseURL string, httpClient *http.Client) API {
	return &api{
		Client:  httpClient,
		baseURL: baseURL,
	}
}

// NoResultsErr is error returned from making an http request to GoData
type NoResultsErr struct {
	Err error
	Msg string
}

func (e *NoResultsErr) Error() string {
	return e.Msg
}

// Unwrap unwraps the error
func (e *NoResultsErr) Unwrap() error {
	return e.Err
}

// DecodeErr happens when decoding an http response fails
type DecodeErr struct {
	Err error
	Msg string
}

func (e *DecodeErr) Error() string {
	return e.Msg
}

// Unwrap unwraps the error
func (e *DecodeErr) Unwrap() error {
	return e.Err
}

// HTTPRequestErr is any http error
type HTTPRequestErr struct {
	Err error
	Msg string
}

func (e *HTTPRequestErr) Error() string {
	return e.Msg
}

// Unwrap unwraps the HTTP error
func (e *HTTPRequestErr) Unwrap() error {
	return e.Err
}

// GetCaseByVisualId retrieves a case from GoData that matches the visualId.
// An error is returned if the http request fails or if no case is found.
func (a *api) GetCaseByVisualId(visualID string, opts Options) (CaseID, error) {
	token := opts.Token
	// We need the id, so we should query for it.
	filter := fmt.Sprintf("{\"where\":{\"visualId\":{\"regexp\":\"/^%s/i\"}}}", visualID)
	getURL := fmt.Sprintf("%s/outbreaks/%s/cases?filter=%s&access_token=%s", a.baseURL, opts.OutbreakID, url.QueryEscape(filter), opts.Token)
	getReq, _ := http.NewRequest(http.MethodGet, getURL, nil)
	getReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	getReq.Header.Set("Content-Type", "application/json")
	getResp, err := a.Client.Do(getReq)
	if err != nil {
		return CaseID{}, &HTTPRequestErr{
			Err: err,
			Msg: "could not retrieve case from GoData",
		}
	}
	defer getResp.Body.Close() //nolint:errcheck
	log.WithFields(log.Fields{
		"body":            getResp.Body,
		"status":          getResp.Status,
		"visualIdRequest": visualID,
		"url":             getURL,
	}).Info("retrieved visualID")

	var resps []CaseID
	if decodeErr := json.NewDecoder(getResp.Body).Decode(&resps); decodeErr != nil {
		log.WithFields(log.Fields{
			"body":   getResp.Body,
			"status": getResp.Status,
		}).WithError(decodeErr).Info("godata case raw body")

		return CaseID{}, &DecodeErr{
			Err: decodeErr,
			Msg: "failed to decode case data",
		}
	}
	if len(resps) == 0 {
		return CaseID{}, &NoResultsErr{
			Err: err,
			Msg: fmt.Sprintf("no record found with visualID: %s", visualID),
		}
	}
	return resps[0], nil
}

func newCase(args CaseArg) GoDataCase {
	screening := args.Screening
	personalInfo := args.PersonalInfo
	arrivalInfo := args.ArrivalInfo
	visualID := args.VisualID
	address := eventToGoDataAddress(args.Address, personalInfo, arrivalInfo.DateOfArrival)
	var caseType = "Non-Tourist"
	if strings.ToLower(arrivalInfo.PurposeOfTrip) == "tourist" {
		caseType = "Tourist"
	}

	goDataQuestionnaire := createGoDataQuestionnaire(personalInfo, screening, arrivalInfo, caseType)
	return createGoDataCase(goDataQuestionnaire, []GoDataAddress{address}, personalInfo, visualID)
}

// UpdateCase makes a put request to GoData to update an existing case.
func (a *api) UpdateCase(args CaseArg, caseID string, opts Options) error {
	godataCase := newCase(args)
	log.WithFields(log.Fields{
		"godataCase": godataCase,
	}).Info("putting to godata")

	return a.putCase(godataCase, caseID, opts.Token)
}

// CreateCase creates a new godata case by making an http post request
func (a *api) CreateCase(args CaseArg, opts Options) error {
	godataCase := newCase(args)
	log.WithFields(log.Fields{
		"godataCase": godataCase,
	}).Info("Posting to godata")

	return postCase(godataCase, opts) // push visitor log struct to
}

// GetGodataToken retrieves a JWT token from a GoData Server.
func GetGodataToken(username, password, baseURL string) (string, error) {
	reqBody, err := json.Marshal(map[string]string{"username": username, "password": password})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error authenticating with GoData")
		return "", fmt.Errorf("marshalling request for godata token failed: %w", err)
	}
	req, err := http.Post(fmt.Sprintf("%s/oauth/token", baseURL),
		"application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error retrieving token from GoData")
		return "", fmt.Errorf("failed to retrieve token from GoData: %w", err)
	}

	defer req.Body.Close() //nolint:errcheck

	var authResp *goDataAuthResponse
	if err := json.NewDecoder(req.Body).Decode(&authResp); err != nil {
		log.WithFields(log.Fields{"error": err, "response": req}).Error("failed to decode oauth token from godata")
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}
	return authResp.AccessToken, nil
}

func postToGodata(godataCase GoDataCase, opts Options) (*http.Response, error) {

	token := opts.Token
	outbreakID := opts.OutbreakID

	// Prepare post request to create case
	body, err := json.Marshal(godataCase)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "data": godataCase}).Error("failed to marshal go data case")
		return nil, fmt.Errorf("failed to marshal godata case: %w", err)
	}

	client := &http.Client{}
	newReq, _ := http.NewRequest("POST", fmt.Sprintf("%s/outbreaks/%s/cases", opts.URL, outbreakID), bytes.NewReader(body))
	newReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	newReq.Header.Set("Content-Type", "application/json")
	log.WithFields(log.Fields{
		"body": body,
		"case": godataCase,
	}).Info("Sending new request to GoData")
	defer newReq.Body.Close() //nolint:errcheck
	return client.Do(newReq)  //nolint:wrapcheck
}

// postCase creates a new case in a GoData server.
func postCase(o GoDataCase, opts Options) error {
	resp, err := postToGodata(o, opts)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to post new case to godata")
		return fmt.Errorf("failed to post new case to godata: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"respBody": string(respBody),
		}).WithError(err).Error("Error reading response from godata")
		return fmt.Errorf("error reading godata response: %v  error: %w", respBody, err)
	}
	log.WithFields(log.Fields{
		"case":     o,
		"response": respBody,
	}).Info("Got a response from godata")

	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{"responseFromGoData": respBody}).Error("error posting case to godata")
		return fmt.Errorf("error posting case to godata") //nolint:goerr113
	}

	log.WithFields(
		log.Fields{
			"case":     o,
			"response": respBody,
		}).Info("posted new case to godata")
	return nil
}

// putCase updates a case in GoData.
func (a *api) putCase(o GoDataCase, caseID, token string) error {
	body, err := json.Marshal(o)
	if err != nil {
		return fmt.Errorf("failed to marshal godata case: %w", err)
	}

	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", a.baseURL, caseID), bytes.NewReader(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.Client.Do(req)
	if err != nil {
		return fmt.Errorf("putCase() to GoData failed: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck
	log.WithFields(log.Fields{
		"case":   o,
		"caseID": caseID,
	}).Info("updated godata case")
	return nil
}

func eventToGoDataAddress(bzAddress models.AddressInBelize, personalInfo models.PersonalInfo, date time.Time) GoDataAddress {
	usualPlaceOfResidence := "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_USUAL_PLACE_OF_RESIDENCE"
	//accommodationResidence := "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_ACCOMMODATION_NAME"

	addressLine1 := bzAddress.Address.Address
	if len(bzAddress.AccommodationName) > 0 {
		addressLine1 = bzAddress.AccommodationName
	}

	address := GoDataAddress{
		TypeId:       usualPlaceOfResidence,
		Country:      "Belize",
		City:         bzAddress.Address.Community.Municipality,
		AddressLine1: addressLine1,
		AddressLine2: bzAddress.Address.Address,
		Date:         date.Format("2006-01-02"),
		PhoneNumber:  personalInfo.PhoneNumbers,
		LocationId:   bzAddress.Address.Community.ID,
		Email:        personalInfo.Email,
	}
	return address
}

func createGoDataQuestionnaire(personalInfo models.PersonalInfo, screening models.Screening, arrivalInfo models.ArrivalInfo, caseType string) GoDataQuestionnaire {
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
	if strings.ToLower(arrivalInfo.PurposeOfTrip) == "tourist" {
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
		PriorXdayExposureContactWithCase:     nil,
		PriorXDayexposureContactWithCaseDate: nil,
		PriorXdayExposureInternationalDateTravelFrom: []PriorXdayExposureInternationalDateTravelFrom{
			{Value: arrivalInfo.DateOfEmbarkation},
		},
		PriorXdayExposureInternationalDatetravelTo: nil,
		PriorXdayexposureInternationaltravelcountries: []PriorXdayExposureInternationalTravelCountries{
			{Value: arrivalInfo.CountriesVisited},
		},
		TypeOfTraveller: []QuestionnaireAnswer{
			{Value: typeOfTraveller},
		},
		PurposeOfTravel: []QuestionnaireAnswer{
			{Value: arrivalInfo.PurposeOfTrip},
		},
		FlightNumber: []QuestionnaireAnswer{
			{Value: arrivalInfo.VesselNumber},
		},
		PriorXdayExposureInternationalTravelCities: nil,
		PcrTestInPast72Hours: []QuestionnaireAnswer{
			{Value: pcrTest},
		},
		PortOfEntry: []QuestionnaireAnswer{{Value: arrivalInfo.PortOfEntry}},
		Nationality: []QuestionnaireAnswer{{Value: personalInfo.Nationality}},
	}

	return godataQuestionnaire
}

func createGoDataCase(goDataQuestionnaire GoDataQuestionnaire, addresses []GoDataAddress, personalInfo models.PersonalInfo, visualId string) GoDataCase {
	gender := "LNG_REFERENCE_DATA_CATEGORY_GENDER_MALE"
	if personalInfo.Gender == "Female" {
		gender = "LNG_REFERENCE_DATA_CATEGORY_GENDER_FEMALE"
	}
	return GoDataCase{
		FirstName:                    personalInfo.FirstName,
		MiddleName:                   personalInfo.MiddleName,
		LastName:                     personalInfo.LastName,
		Gender:                       gender,
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
