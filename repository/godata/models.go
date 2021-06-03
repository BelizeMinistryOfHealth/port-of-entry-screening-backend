package godata

import "time"

// GoDataAddress
type GoDataAddress struct {
	TypeId       string `json:"typeId"`
	Country      string `json:"country"`
	City         string `json:"city"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Date         string `json:"date"`
	PhoneNumber  string `json:"phoneNumber"`
	LocationId   string `json:"locationId"`
}

type GeoLocation struct {
	Lat int32 `json:"lat"`
	Lng int32 `json:"lng"`
}

type GoDataCaseForm struct {
	Value []string `json:"value"`
}

type DataCollectorName struct {
	Value string `json:"value"`
}
type CaseCountryResidence struct {
	Value string `json:"value"`
}
type ShowsSymptoms struct {
	Value string `json:"value"`
}
type SymptomFever struct {
	Value string `json:"value"`
}
type SoreThroat struct {
	Value string `json:"value"`
}
type RunnyNose struct {
	Value string `json:"value"`
}
type Cough struct {
	Value string `json:"value"`
}
type Vomiting struct {
	Value string `json:"value"`
}
type Nausea struct {
	Value string `json:"value"`
}
type Diarrhea struct {
	Value string `json:"value"`
}
type ShortnessOfBreath struct {
	Value string `json:"value"`
}

type DifficultyBreathing struct {
	Value string `json:"value"`
}

type SymptomsChills struct {
	Value string `json:"value"`
}

type Headache struct {
	Value string `json:"value"`
}

type Malaise struct {
	Value string `json:"value"`
}

type Anosmia struct {
	Value string `json:"value"`
}

type Aguesia struct {
	Value string `json:"value"`
}

type Bleeding struct {
	Value string `json:"value"`
}

type JointMusclePain struct {
	Value string `json:"value"`
}

type EyeFacialPain struct {
	Value string `json:"value"`
}

type GeneralizedRash struct {
	Value string `json:"value"`
}

type BlurredVision struct {
	Value string `json:"value"`
}

type AbdominalPain struct {
	Value string `json:"value"`
}

type PriorXdayExposureTravelledInternationally struct {
	Value string `json:"value"`
}
type PriorXdayExposureContactWithCase struct {
	Value string `json:"value"`
}
type PriorXdayExposureContactWithCaseDate struct {
	Value time.Time `json:"value"`
}
type PriorXdayExposureInternationalDateTravelFrom struct {
	Value time.Time `json:"value"`
}
type PriorXdayExposureInternationalDateTravelTo struct {
	Value time.Time `json:"value"`
}
type PriorXdayExposureInternationalTravelCountries struct {
	Value string `json:"value"`
}
type PriorXdayExposureInternationalTravelCities struct {
	Value string `json:"value"`
}

type QuestionnaireAnswer struct {
	Value string `json:"value"`
}

type GoDataQuestionnaire struct {
	CaseForm                                      []GoDataCaseForm                                `json:"Case_WhichForm"`
	DataCollectorName                             []DataCollectorName                             `json:"FA0_datacollector_name"`
	CountryResidence                              []CaseCountryResidence                          `json:"FA0_case_countryresidence"`
	ShowsSymptoms                                 []ShowsSymptoms                                 `json:"FA0_symptoms_caseshowssymptoms"`
	Fever                                         []SymptomFever                                  `json:"FA0_symptom_fever"`
	SoreThroat                                    []SoreThroat                                    `json:"FA0_symptom_sorethroat"`
	RunnyNose                                     []RunnyNose                                     `json:"FA0_symptom_runnynose"`
	Cough                                         []Cough                                         `json:"FA0_symptom_cough"`
	Vomiting                                      []Vomiting                                      `json:"FA0_symptom_vomiting"`
	Nausea                                        []Nausea                                        `json:"FA0_symptom_nausea"`
	Diarrhea                                      []Diarrhea                                      `json:"FA0_symptom_diarrhea"`
	ShortnessOfBreath                             []ShortnessOfBreath                             `json:"FA0_symptom_shortnessofbreath"`
	DifficultyBreathing                           []DifficultyBreathing                           `json:"FA0_symptom_difficulty_breathing"`
	SymptomsChills                                []SymptomsChills                                `json:"FA0_symptom_chills"`
	Headache                                      []Headache                                      `json:"FA0_symptom_headache"`
	Malaise                                       []Malaise                                       `json:"FA0_symptom_malaise"`
	Anosmia                                       []Anosmia                                       `json:"FA0_symptom_anosmia"`
	Aguesia                                       []Aguesia                                       `json:"FA0_symptom_aguesia"`
	Bleeding                                      []Bleeding                                      `json:"FA0_symptom_bleeding"`
	JointMusclePain                               []JointMusclePain                               `json:"FA0_symptom_joint_muscle_pain"`
	EyeFacialPain                                 []EyeFacialPain                                 `json:"FA0_symptom_eye_facial_pain"`
	GeneralizedRash                               []GeneralizedRash                               `json:"FA0_symptom_generalized_rash"`
	BlurredVision                                 []BlurredVision                                 `json:"FA0_symptom_blurred_vision"`
	AbdominalPain                                 []AbdominalPain                                 `json:"FA0_symptom_abdominal_pain"`
	CaseType                                      string                                          `json:"case_type"`
	PriorXdayExposureTravelledInternationally     []QuestionnaireAnswer                           `json:"FA0_priorXdayexposure_travelledinternationally"`
	PriorXdayExposureContactWithCase              []PriorXdayExposureContactWithCase              `json:"FA0_priorXdayexposure_contactwithcase"`
	PriorXDayexposureContactWithCaseDate          []PriorXdayExposureContactWithCaseDate          `json:"FA0_priorXdayexposure_contactwithcasedate"`
	PriorXdayExposureInternationalDateTravelFrom  []PriorXdayExposureInternationalDateTravelFrom  `json:"FA0_priorXdayexposure_internationaldatetravelfrom"`
	PriorXdayExposureInternationalDatetravelTo    []PriorXdayExposureInternationalDateTravelTo    `json:"FA0_priorXdayexposure_internationaldatetravelto"`
	PriorXdayexposureInternationaltravelcountries []PriorXdayExposureInternationalTravelCountries `json:"FA0_priorXdayexposure_internationaltravelcountries"`
	PriorXdayExposureInternationalTravelCities    []PriorXdayExposureInternationalTravelCities    `json:"FA0_priorXdayexposure_internationaltravelcities"`
	TypeOfTraveller                               []QuestionnaireAnswer                           `json:"FA0_priorXdayexposure_typeoftraveler"`
	PurposeOfTravel                               []QuestionnaireAnswer                           `json:"FA0_priorXdayexposure_purposeoftravel"`
	FlightNumber                                  []QuestionnaireAnswer                           `json:"FA0_priorXdayexposure_flightnumber"`
	PcrTestInPast72Hours                          []QuestionnaireAnswer                           `json:"FA0_priorXdayexposure_tookpcrtest_past72hours"`
	PortOfEntry                                   []QuestionnaireAnswer                           `json:"port_of_entry"`
}

type GoDataCase struct {
	FirstName                    string              `json:"firstName"`
	MiddleName                   string              `json:"middleName"`
	LastName                     string              `json:"lastName"`
	Gender                       string              `json:"gender"`
	Classification               string              `json:"classification"`
	Dob                          string              `json:"dob"`
	Occupation                   string              `json:"occupation"`
	DateOfReporting              string              `json:"dateOfReporting"`
	IsDateOfReportingApproximate bool                `json:"isDateOfReportingApproximate"`
	Addresses                    []GoDataAddress     `json:"addresses"`
	Questionnaire                GoDataQuestionnaire `json:"questionnaireAnswers"`
	VisualId                     string              `json:"visualId"`
}

type GoDataPerson struct {
	Id                           string              `json:"id"`
	FirstName                    string              `json:"firstName"`
	MiddleName                   string              `json:"middleName"`
	LastName                     string              `json:"lastName"`
	Gender                       string              `json:"gender"`
	Classification               string              `json:"classification"`
	Dob                          string              `json:"dob"`
	Occupation                   string              `json:"occupation"`
	DateOfReporting              string              `json:"dateOfReporting"`
	IsDateOfReportingApproximate bool                `json:"isDateOfReportingApproximate"`
	Addresses                    []GoDataAddress     `json:"addresses"`
	Questionnaire                GoDataQuestionnaire `json:"questionnaireAnswers"`
	VisualId                     string              `json:"visualId"`
}
