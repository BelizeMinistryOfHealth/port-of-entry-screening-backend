package godata

import "time"

// Address is the address representation in GoData
type Address struct {
	TypeID       string `json:"typeId"`
	Country      string `json:"country"`
	City         string `json:"city"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Date         string `json:"date"`
	PhoneNumber  string `json:"phoneNumber"`
	LocationID   string `json:"locationId"`
	Email        string `json:"emailAddress"`
}

// GeoLocation is a geo representation
type GeoLocation struct {
	Lat int32 `json:"lat"`
	Lng int32 `json:"lng"`
}

// QuestionnaireTimeAnswer is an date answer
type QuestionnaireTimeAnswer struct {
	Value time.Time `json:"value"`
}

// QuestionnaireAnswer is an answer with a string
type QuestionnaireAnswer struct {
	Value string `json:"value"`
}

// CaseForm is an answer with a []string
type CaseForm struct {
	Value []string `json:"value"`
}

// Questionnaire is the representation of the godata questionnaire
type Questionnaire struct {
	CaseForm                                      []CaseForm                `json:"Case_WhichForm"`
	DataCollectorName                             []QuestionnaireAnswer     `json:"FA0_datacollector_name"`
	CountryResidence                              []QuestionnaireAnswer     `json:"FA0_case_countryresidence"`
	ShowsSymptoms                                 []QuestionnaireAnswer     `json:"FA0_symptoms_caseshowssymptoms"`
	Fever                                         []QuestionnaireAnswer     `json:"FA0_symptom_fever"`
	SoreThroat                                    []QuestionnaireAnswer     `json:"FA0_symptom_sorethroat"`
	RunnyNose                                     []QuestionnaireAnswer     `json:"FA0_symptom_runnynose"`
	Cough                                         []QuestionnaireAnswer     `json:"FA0_symptom_cough"`
	Vomiting                                      []QuestionnaireAnswer     `json:"FA0_symptom_vomiting"`
	Nausea                                        []QuestionnaireAnswer     `json:"FA0_symptom_nausea"`
	Diarrhea                                      []QuestionnaireAnswer     `json:"FA0_symptom_diarrhea"`
	ShortnessOfBreath                             []QuestionnaireAnswer     `json:"FA0_symptom_shortnessofbreath"`
	DifficultyBreathing                           []QuestionnaireAnswer     `json:"FA0_symptom_difficulty_breathing"`
	SymptomsChills                                []QuestionnaireAnswer     `json:"FA0_symptom_chills"`
	Headache                                      []QuestionnaireAnswer     `json:"FA0_symptom_headache"`
	Malaise                                       []QuestionnaireAnswer     `json:"FA0_symptom_malaise"`
	Anosmia                                       []QuestionnaireAnswer     `json:"FA0_symptom_anosmia"`
	Aguesia                                       []QuestionnaireAnswer     `json:"FA0_symptom_aguesia"`
	Bleeding                                      []QuestionnaireAnswer     `json:"FA0_symptom_bleeding"`
	JointMusclePain                               []QuestionnaireAnswer     `json:"FA0_symptom_joint_muscle_pain"`
	EyeFacialPain                                 []QuestionnaireAnswer     `json:"FA0_symptom_eye_facial_pain"`
	GeneralizedRash                               []QuestionnaireAnswer     `json:"FA0_symptom_generalized_rash"`
	BlurredVision                                 []QuestionnaireAnswer     `json:"FA0_symptom_blurred_vision"`
	AbdominalPain                                 []QuestionnaireAnswer     `json:"FA0_symptom_abdominal_pain"`
	CaseType                                      string                    `json:"case_type"`
	PriorXdayExposureTravelledInternationally     []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_travelledinternationally"`
	PriorXdayExposureContactWithCase              []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_contactwithcase"`
	PriorXDayexposureContactWithCaseDate          []QuestionnaireTimeAnswer `json:"FA0_priorXdayexposure_contactwithcasedate"`
	PriorXdayExposureInternationalDateTravelFrom  []QuestionnaireTimeAnswer `json:"FA0_priorXdayexposure_internationaldatetravelfrom"`
	PriorXdayExposureInternationalDatetravelTo    []QuestionnaireTimeAnswer `json:"FA0_priorXdayexposure_internationaldatetravelto"`
	PriorXdayexposureInternationaltravelcountries []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_internationaltravelcountries"`
	PriorXdayExposureInternationalTravelCities    []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_internationaltravelcities"`
	TypeOfTraveller                               []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_typeoftraveler"`
	PurposeOfTravel                               []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_purposeoftravel"`
	FlightNumber                                  []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_flightnumber"`
	PcrTestInPast72Hours                          []QuestionnaireAnswer     `json:"FA0_priorXdayexposure_tookpcrtest_past72hours"`
	PortOfEntry                                   []QuestionnaireAnswer     `json:"port_of_entry"`
	Nationality                                   []QuestionnaireAnswer     `json:"nationality"`
}

// Case is a representation of a godata case
type Case struct {
	FirstName                    string        `json:"firstName"`
	MiddleName                   string        `json:"middleName"`
	LastName                     string        `json:"lastName"`
	Gender                       string        `json:"gender"`
	Classification               string        `json:"classification"`
	Dob                          string        `json:"dob"`
	Occupation                   string        `json:"occupation"`
	DateOfReporting              string        `json:"dateOfReporting"`
	IsDateOfReportingApproximate bool          `json:"isDateOfReportingApproximate"`
	Addresses                    []Address     `json:"addresses"`
	Questionnaire                Questionnaire `json:"questionnaireAnswers"`
	VisualID                     string        `json:"visualId"`
	Vaccines                     []Vaccination `json:"vaccinesReceived"`
}

// Person is a representation of a person in GoData
type Person struct {
	ID                           string        `json:"id"`
	FirstName                    string        `json:"firstName"`
	MiddleName                   string        `json:"middleName"`
	LastName                     string        `json:"lastName"`
	Gender                       string        `json:"gender"`
	Classification               string        `json:"classification"`
	Dob                          string        `json:"dob"`
	Occupation                   string        `json:"occupation"`
	DateOfReporting              string        `json:"dateOfReporting"`
	IsDateOfReportingApproximate bool          `json:"isDateOfReportingApproximate"`
	Addresses                    []Address     `json:"addresses"`
	Questionnaire                Questionnaire `json:"questionnaireAnswers"`
	VisualID                     string        `json:"visualId"`
}

// VaccineStatus is a person's COVID19 vaccination status
type VaccineStatus string

// VaccineDose is the COVID19 vaccine a person received
type VaccineDose string

const (
	// Vaccinated indicates that a person is vaccinated
	Vaccinated VaccineStatus = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_STATUS_VACCINATED"
	// NotVaccinated indicates that a person is not vaccinated
	NotVaccinated VaccineStatus = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_STATUS_NOT_VACCINATED"
	// AstraFirst is the first dose of AstraZeneca
	AstraFirst VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_ASTRAZENECA_1ST_DOSE"
	// AstraSecond is the second dose of AstraZeneca
	AstraSecond   VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_ASTRAZENECA_2ND_DOSE"
	Johnson       VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_JOHNSON_JOHNSON"
	ModernaFirst  VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_MODERNA_1ST_DOSE"
	ModernaSecond VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_MODERNA_2ND_DOSE"
	PfizerFirst   VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_PFIZER_1ST_DOSE"
	PfizerSecond  VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_PFIZER_2ND_DOSE"
	SinoFirst     VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_SINOPHARM_1ST_DOSE"
	SinoSecond    VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_SINOPHARM_2ND_DOSE"
	SputnikFirst  VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_SPUTNIK_1ST_DOSE"
	SputnikSecond VaccineDose = "LNG_REFERENCE_DATA_CATEGORY_VACCINE_SPUTNIK_2ND_DOSE"
)

// Vaccination structure in Godata is an array:
// vaccinesReceived: [{date, status (LNG_REFERENCE_DATA_CATEGORY_VACCINE_STATUS_VACCINATED), vaccine (LNG_REFERENCE_DATA_CATEGORY_VACCINE_ASTRAZENECA_1ST_DOSE)}]
type Vaccination struct {
	Date    time.Time     `json:"date"`
	Status  VaccineStatus `json:"status"`
	Vaccine VaccineDose   `json:"vaccine"`
}
