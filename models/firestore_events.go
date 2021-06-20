package models

import (
	"strconv"
	"time"
)

// FirestoreAddresses is how the addresses are encoded in a Firestore event
type FirestoreAddresses struct {
	ArrayValue struct {
		Values []struct {
			MapValue struct {
				Fields struct {
					ControlID struct {
						StringValue string `json:"stringValue"`
					} `json:"controlId"`
					AccommodationName struct {
						StringValue string `json:"stringValue"`
					} `json:"accommodationName"`
					EndDate struct {
						StringValue string `json:"stringValue"`
					} `json:"endDate"`
					StartDate struct {
						StringValue string `json:"stringValue"`
					} `json:"startDate"`
					Address struct {
						MapValue struct {
							Fields struct {
								Address struct {
									StringValue string `json:"stringValue"`
								} `json:"address"`
								Community struct {
									MapValue struct {
										Fields struct {
											District struct {
												StringValue string `json:"stringValue"`
											} `json:"district"`
											ID struct {
												StringValue string `json:"stringValue"`
											} `json:"id"`
											Municipality struct {
												StringValue string `json:"stringValue"`
											} `json:"municipality"`
										} `json:"fields"`
									} `json:"mapValue"`
								} `json:"community"`
							} `json:"fields"`
						} `json:"mapValue"`
					} `json:"address"`
				} `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}

// FluLikeSymptomsEvent represents the flu symptoms in a firestore event
type FluLikeSymptomsEvent struct {
	MapValue struct {
		Fields struct {
			AbdominalPain struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"abdominalPain"`
			Aguesia struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"aguesia"`
			Anosmia struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"anosmia"`
			Bleeding struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"bleeding"`
			BlurredVision struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"blurredVision"`
			BreathDifficulty struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"breathDifficulty"`
			BreathShort struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"breathShort"`
			Chills struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"chills"`
			Cough struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"cough"`
			Diarrhea struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"diarrhea"`
			EyeFacialPain struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"eyeFacialPain"`
			Fever struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"fever"`
			GeneralizedRash struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"generalizedRash"`
			Headache struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"headache"`
			JointMusclePain struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"jointMusclePain"`
			Malaise struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"malaise"`
			Nausea struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"nausea"`
			Other struct {
				StringValue string `json:"stringValue"`
			} `json:"other"`
			RunnyNose struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"runnyNose"`
			ShortnessOfBreath struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"shortnessOfBreath"`
			SoreThroat struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"soreThroat"`
			Vomiting struct {
				BooleanValue bool `json:"booleanValue"`
			} `json:"vomiting"`
		} `json:"fields"`
	} `json:"mapValue"`
}

// ToFluLikeSymptoms converts an event structure to the FluLikeSymptoms domain structure.
func (f *FluLikeSymptomsEvent) ToFluLikeSymptoms() FluLikeSymptoms {
	return FluLikeSymptoms{
		Fever:             f.MapValue.Fields.Fever.BooleanValue,
		Headache:          f.MapValue.Fields.Headache.BooleanValue,
		Cough:             f.MapValue.Fields.Cough.BooleanValue,
		Malaise:           f.MapValue.Fields.Malaise.BooleanValue,
		SoreThroat:        f.MapValue.Fields.SoreThroat.BooleanValue,
		BreathShort:       f.MapValue.Fields.BreathShort.BooleanValue,
		BreathDifficulty:  f.MapValue.Fields.BreathDifficulty.BooleanValue,
		RunnyNose:         f.MapValue.Fields.RunnyNose.BooleanValue,
		Nausea:            f.MapValue.Fields.Nausea.BooleanValue,
		Diarrhea:          f.MapValue.Fields.Diarrhea.BooleanValue,
		ShortnessOfBreath: f.MapValue.Fields.ShortnessOfBreath.BooleanValue,
		Chills:            f.MapValue.Fields.Chills.BooleanValue,
		Anosmia:           f.MapValue.Fields.Anosmia.BooleanValue,
		Aguesia:           f.MapValue.Fields.Aguesia.BooleanValue,
		Bleeding:          f.MapValue.Fields.Bleeding.BooleanValue,
		JointMusclePain:   f.MapValue.Fields.JointMusclePain.BooleanValue,
		EyeFacialPain:     f.MapValue.Fields.EyeFacialPain.BooleanValue,
		GeneralizedRash:   f.MapValue.Fields.GeneralizedRash.BooleanValue,
		BlurredVision:     f.MapValue.Fields.BlurredVision.BooleanValue,
		AbdominalPain:     f.MapValue.Fields.AbdominalPain.BooleanValue,
		Vomiting:          f.MapValue.Fields.Vomiting.BooleanValue,
		Other:             f.MapValue.Fields.Other.StringValue,
	}
}

// VaccinationEvent represents the vaccination value in a Firestore event
type VaccinationEvent struct {
	MapValue struct {
		Fields struct {
			DateOfMostRecentShot struct {
				TimestampValue string `json:"timestampValue"`
			} `json:"dateOfMostRecentShot"`
			Name struct {
				StringValue string `json:"stringValue"`
			} `json:"name"`
			NumberOfShots struct {
				IntegerValue string `json:"integerValue"`
			} `json:"numberOfShots"`
		} `json:"fields"`
	} `json:"mapValue"`
}

// ToVaccination converts a vaccination representation in an event to a Vaccination domain model
func (v *VaccinationEvent) ToVaccination() Vaccination {
	shots, err := strconv.Atoi(v.MapValue.Fields.NumberOfShots.IntegerValue)
	if err != nil {
		shots = 0
	}
	dateShot, err := time.Parse("2006-01-02", v.MapValue.Fields.DateOfMostRecentShot.TimestampValue)
	if err != nil {
		dateShot = time.Now()
	}
	return Vaccination{
		Name:                 v.MapValue.Fields.Name.StringValue,
		NumberOfShots:        shots,
		DateOfMostRecentShot: dateShot,
	}
}

// EditorEvent represents the editor's data in a firestore event
type EditorEvent struct {
	MapValue struct {
		Fields struct {
			Email struct {
				StringValue string `json:"stringValue"`
			} `json:"email"`
			ID struct {
				StringValue string `json:"stringValue"`
			} `json:"id"`
		} `json:"fields"`
	} `json:"mapValue"`
}

// ToEditor converts the editor structure in a firestore event to the domain Editor model
func (e *EditorEvent) ToEditor() Editor {
	return Editor{
		Email: e.MapValue.Fields.Email.StringValue,
		ID:    e.MapValue.Fields.ID.StringValue,
	}
}

// FirestoreScreenings is how the screenings are represented in a Firestore event
type FirestoreScreenings struct {
	Comments struct {
		StringValue string `json:"stringValue"`
	} `json:"comments"`
	CreatedBy            EditorEvent `json:"createdBy"`
	DiagnosedWithCovid19 struct {
		BooleanValue bool `json:"booleanValue"`
	} `json:"diagnosedWithCovid19"`
	FluLikeSymptoms FluLikeSymptomsEvent `json:"fluLikeSymptoms"`
	ID              struct {
		StringValue string `json:"stringValue"`
	} `json:"id"`
	Location struct {
		StringValue string `json:"stringValue"`
	} `json:"location"`
	Modified struct {
		TimestampValue string `json:"timestampValue"`
	} `json:"modified"`
	ModifiedBy EditorEvent `json:"modifiedBy"`
	Screened   struct {
		TimestampValue string `json:"timestampValue"`
	} `json:"screened"`
	Temperature struct {
		IntegerValue string `json:"integerValue"`
	} `json:"temperature"`
	TookPcrTestInPast72Hours struct {
		BooleanValue bool `json:"booleanValue"`
	} `json:"tookPcrTestInPast72Hours"`
	Vaccination VaccinationEvent `json:"vaccination"`
}

// ToScreening converts the screening value in a firestore event to the Screening domain model
func (f *FirestoreScreenings) ToScreening() Screening {
	dateScreened, err := time.Parse("2006-01-02", f.Screened.TimestampValue)
	if err != nil {
		dateScreened = time.Now()
	}
	temperature, tempErr := strconv.ParseFloat(f.Temperature.IntegerValue, 32)
	if tempErr != nil {
		temperature = 0.0
	}
	return Screening{
		ID:                       f.ID.StringValue,
		DiagnosedWithCovid19:     f.DiagnosedWithCovid19.BooleanValue,
		Comments:                 f.Comments.StringValue,
		Location:                 f.Location.StringValue,
		DateScreened:             dateScreened,
		Temperature:              float32(temperature),
		FluLikeSymptoms:          f.FluLikeSymptoms.ToFluLikeSymptoms(),
		TookPcrTestInPast72Hours: f.TookPcrTestInPast72Hours.BooleanValue,
		Vaccination:              f.Vaccination.ToVaccination(),
		Modified:                 nil,
		CreatedBy:                Editor{},
		ModifiedBy:               Editor{},
	}
}

// StringValueStruct represents a string value in a Firestore Event
type StringValueStruct struct {
	StringValue string `json:"stringValue"`
}

// IntValueStruct represents an int value in a Firestore Event
type IntValueStruct struct {
	IntValue string `json:"intValue"`
}

// TimestampValueStruct represents a time value in a Firestore event
type TimestampValueStruct struct {
	TimestampValue time.Time `json:"timestampValue"`
}

// PersonFirestoreFields represents the person value in a firestore event
type PersonFirestoreFields struct {
	ID                    StringValueStruct `json:"id"`
	FirstName             StringValueStruct `json:"firstName"`
	MiddleName            StringValueStruct `json:"middleName"`
	LastName              StringValueStruct `json:"lastName"`
	Gender                StringValueStruct `json:"gender"`
	FullName              StringValueStruct `json:"fullName"`
	Dob                   StringValueStruct `json:"dob"`
	Nationality           StringValueStruct `json:"nationality"`
	PhoneNumbers          StringValueStruct `json:"phoneNumbers"`
	PassportNumber        StringValueStruct `json:"passportNumber"`
	OtherTravelDocument   StringValueStruct `json:"otherTravelDocument"`
	OtherTravelDocumentID StringValueStruct `json:"otherTravelDocumentId"`
	Email                 StringValueStruct `json:"email"`
	BhisNumber            StringValueStruct `json:"bhisNumber"`
	PortOfEntry           StringValueStruct `json:"portOfEntry"`
	Occupation            StringValueStruct `json:"occupation"`
	CreatedBy             struct {
		MapValueStruct struct {
			Fields struct {
				ID    StringValueStruct `json:"id"`
				Email StringValueStruct `json:"email"`
			} `json:"Fields"`
		} `json:"mapValue"`
	} `json:"createdBy"`
	Created    TimestampValueStruct `json:"created"`
	ModifiedBy struct {
		MapValueStruct struct {
			Fields struct {
				ID    StringValueStruct `json:"id"`
				Email StringValueStruct `json:"email"`
			} `json:"fields"`
		} `json:"mapValue"`
	} `json:"modifiedBy"`
	Modified TimestampValueStruct `json:"modified"`
}

// ToPerson converts the person value in a firestore event to a PersonalInfo domain model
func (p *PersonFirestoreFields) ToPerson() PersonalInfo {
	return PersonalInfo{
		ID:                    p.ID.StringValue,
		FirstName:             p.FirstName.StringValue,
		LastName:              p.LastName.StringValue,
		MiddleName:            p.MiddleName.StringValue,
		FullName:              p.FullName.StringValue,
		Dob:                   p.Dob.StringValue,
		Nationality:           p.Nationality.StringValue,
		PassportNumber:        p.PassportNumber.StringValue,
		OtherTravelDocument:   p.OtherTravelDocument.StringValue,
		OtherTravelDocumentID: p.OtherTravelDocumentID.StringValue,
		Email:                 p.Email.StringValue,
		Gender:                p.Gender.StringValue,
		PhoneNumbers:          p.PhoneNumbers.StringValue,
		BhisNumber:            p.BhisNumber.StringValue,
		Occupation:            p.Occupation.StringValue,
		PortOfEntry:           p.PortOfEntry.StringValue,
		Created:               p.Created.TimestampValue,
		Modified:              p.Modified.TimestampValue,
		CreatedBy: Editor{
			ID:    p.CreatedBy.MapValueStruct.Fields.ID.StringValue,
			Email: p.CreatedBy.MapValueStruct.Fields.Email.StringValue,
		},
		ModifiedBy: Editor{
			ID:    p.ModifiedBy.MapValueStruct.Fields.ID.StringValue,
			Email: p.ModifiedBy.MapValueStruct.Fields.Email.StringValue,
		},
	}
}

// FirestorePersonEvent represents a Firestore Event
type FirestorePersonEvent struct {
	OldValue   FirestorePersonValue `json:"oldValue"`
	Value      FirestorePersonValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestorePersonValue represents the value in a firestore event
type FirestorePersonValue struct {
	CreateTime time.Time             `json:"createTime"`
	Fields     PersonFirestoreFields `json:"fields"`
	Name       string                `json:"name"`
	UpdateTime time.Time             `json:"updateTime"`
}

// FirestoreScreeningEvent represents a Screening Firestore event
type FirestoreScreeningEvent struct {
	OldValue   FirestoreScreeningValue `json:"oldValue"`
	Value      FirestoreScreeningValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreScreeningValue represents a value in a Screening Firestore event
type FirestoreScreeningValue struct {
	CreateTime time.Time           `json:"createTime"`
	Fields     FirestoreScreenings `json:"fields"`
	Name       string              `json:"name"`
	UpdateTime time.Time           `json:"updateTime"`
}

// ArrivalFields represents the fields of the arrival record in a firestore event.
type ArrivalFields struct {
	ID                   StringValueStruct    `json:"id"`
	PortOfEntry          StringValueStruct    `json:"portOfEntry"`
	DateOfArrival        TimestampValueStruct `json:"dateOfArrival"`
	CountryOfEmbarkation StringValueStruct    `json:"countryOfEmbarkation"`
}

// FirestoreArrivalValue is the `value` field in a firestore event
type FirestoreArrivalValue struct {
	CreateTime time.Time     `json:"createTime"`
	Fields     ArrivalFields `json:"fields"`
	Name       string        `json:"name"`
	UpdateTime time.Time     `json:"updateTime"`
}

// FirestoreArrivalEvent represents the event triggered from the arrivals firestore collection
type FirestoreArrivalEvent struct {
	OldValue   FirestoreArrivalValue `json:"oldValue"`
	Value      FirestoreArrivalValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}
