package models

import "time"

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

// FirestoreScreenings is how the screenings are represented in a Firestore event
type FirestoreScreenings struct {
	ArrayValue struct {
		Values []struct {
			MapValue struct {
				Fields struct {
					ContactWithHealthFacility struct {
						BooleanValue bool `json:"booleanValue"`
					} `json:"contactWithHealthFacility"`
					OtherSymptoms struct {
						StringValue string `json:"stringValue"`
					} `json:"otherSymptoms"`
					Location struct {
						StringValue string `json:"stringValue"`
					} `json:"location"`
					DiagnosedWithCovid19 struct {
						BooleanValue bool `json:"booleanValue"`
					} `json:"diagnosedWithCovid19"`
					Comments struct {
						StringValue string `json:"stringValue"`
					} `json:"comments"`
					TookPcrTestInPast72Hours struct {
						BooleanValue bool `json:"boolean"`
					} `json:"tookPcrTestInPast72Hours"`
					FluLikeSymptoms struct {
						MapValue struct {
							Fields struct {
								Malaise struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"malaise"`
								Cough struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"cough"`
								Headache struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"headache"`
								BreathDifficulty struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"breathDifficulty"`
								SoreThroat struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"soreThroat"`
								OtherFluLikeSymptoms struct {
									StringValue string `json:"stringValue"`
								} `json:"otherFluLikeSymptoms"`
								BreathShort struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"breathShort"`
								Fever struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"fever"`
								RunnyNose struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"runnyNose"`
								Nausea struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"nausea"`
								Diarrhea struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"diarrhea"`
								ShortnessOfBreath struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"shortnessOfBreath"`
								Chills struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"chills"`
								Anosmia struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"anosmia"`
								Aguesia struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"aguesia"`
								Bleeding struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"bleeding"`
								JointMusclePain struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"jointMusclePain"`
								EyeFacialPain struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"eyeFacialPain"`
								GeneralizedRash struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"generalizedRash"`
								BlurredVision struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"blurredVision"`
								AbdominalPain struct {
									BooleanValue bool `json:"booleanValue"`
								} `json:"abdominalPain"`
								Other struct {
									StringValue string `json:"stringValue"`
								} `json:"other"`
							} `json:"fields"`
						} `json:"mapValue"`
					} `json:"fluLikeSymptoms"`
					Temperature struct {
						DoubleValue float32 `json:"doubleValue"`
					} `json:"temperature"`
					ID struct {
						StringValue string `json:"stringValue"`
					} `json:"id"`
					DateScreened struct {
						StringValue string `json:"stringValue"`
					} `json:"dateScreened"`
				} `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}

// ArrivalFirestoreFields is how the arrivals are represented in a Firestore event
type ArrivalFirestoreFields struct {
	PersonalInfo struct {
		MapValue struct {
			Fields struct {
				PhoneNumbers struct {
					StringValue string `json:"stringValue"`
				} `json:"phoneNumbers"`
				Gender struct {
					StringValue string `json:"stringValue"`
				} `json:"gender"`
				OtherTravelDocument struct {
					StringValue string `json:"stringValue"`
				} `json:"otherTravelDocument"`
				FullName struct {
					StringValue string `json:"stringValue"`
				} `json:"fullName"`
				Dob struct {
					StringValue string `json:"stringValue"`
				} `json:"dob"`
				MiddleName struct {
					StringValue string `json:"stringValue"`
				} `json:"middleName"`
				OtherTravelDocumentID struct {
					StringValue string `json:"stringValue"`
				} `json:"otherTravelDocumentId"`
				LastName struct {
					StringValue string `json:"stringValue"`
				} `json:"lastName"`
				FirstName struct {
					StringValue string `json:"stringValue"`
				} `json:"firstName"`
				PassportNumber struct {
					StringValue string `json:"stringValue"`
				} `json:"passportNumber"`
				Email struct {
					StringValue string `json:"stringValue"`
				} `json:"email"`
				Nationality struct {
					StringValue string `json:"stringValue"`
				} `json:"nationality"`
				BhisNumber struct {
					StringValue string `json:"stringValue"`
				} `json:"bhisNumber"`
				Occupation struct {
					StringValue string `json:"stringValue"`
				} `json:"occupation"`
			} `json:"fields"`
		} `json:"mapValue"`
	} `json:"personalInfo"`
	Arrivals struct {
		MapValue struct {
			Fields struct {
				QuarantineLocation struct {
					StringValue string `json:"stringValue"`
				} `json:"quarantineLocation"`
				Created struct {
					StringValue string `json:"stringValue"`
				} `json:"created"`
				ContactPersonPhoneNumber struct {
					StringValue string `json:"stringValue"`
				} `json:"contactPersonPhoneNumber"`
				ContactPerson struct {
					StringValue string `json:"stringValue"`
				} `json:"contactPerson"`
				ArrivalInfo struct {
					MapValue struct {
						Fields struct {
							CountriesVisited struct {
								StringValue string `json:"stringValue"`
							} `json:"countriesVisited"`
							ModeOfTravel struct {
								StringValue string `json:"stringValue"`
							} `json:"modeOfTravel"`
							PortOfEntry struct {
								StringValue string `json:"stringValue"`
							} `json:"portOfEntry"`
							VesselNumber struct {
								StringValue string `json:"stringValue"`
							} `json:"vesselNumber"`
							DateOfArrival struct {
								StringValue string `json:"stringValue"`
							} `json:"dateOfArrival"`
							TravelOrigin struct {
								StringValue string `json:"stringValue"`
							} `json:"travelOrigin"`
							CountryOfEmbarkation struct {
								StringValue string `json:"stringValue"`
							} `json:"countryOfEmbarkation"`
							DateOfEmbarkation struct {
								StringValue string `json:"stringValue"`
							} `json:"dateOfEmbarkation"`
						} `json:"fields"`
					} `json:"mapValue"`
				} `json:"arrivalInfo"`
				Modified struct {
					StringValue string `json:"stringValue"`
				} `json:"modified"`
				Addresses FirestoreAddresses `json:"addresses"`
				TripID    struct {
					StringValue string `json:"stringValue"`
				} `json:"tripId"`
				PurposeOfTrip struct {
					StringValue string `json:"stringValue"`
				} `json:"purposeOfTrip"`
				Screenings           FirestoreScreenings `json:"screenings"`
				TravellingCompanions struct {
					NullValue interface{} `json:"nullValue"`
				} `json:"travellingCompanions"`
				LengthStay struct {
					StringValue interface{} `json:"nullValue"`
				} `json:"lengthStay"`
			} `json:"fields"`
		} `json:"mapValue"`
	} `json:"arrivals"`
	Modified struct {
		StringValue string `json:"stringValue"`
	} `json:"modified"`
	Created struct {
		StringValue string `json:"stringValue"`
	} `json:"created"`
	PortOfEntry struct {
		StringValue string `json:"stringValue"`
	} `json:"portOfEntry"`
	ID struct {
		StringValue string `json:"stringValue"`
	} `json:"id"`
	ObjectID struct {
		StringValue string `json:"stringValue"`
	} `json:"objectID"`
}

// FirestoreArrivalEvent is the payload of a Firestore event.
type FirestoreArrivalEvent struct {
	OldValue   FirestoreArrivalValue `json:"oldValue"`
	Value      FirestoreArrivalValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreArrivalValue holds Firestore fields.
type FirestoreArrivalValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log an interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     ArrivalFirestoreFields `json:"fields"`
	Name       string                 `json:"name"`
	UpdateTime time.Time              `json:"updateTime"`
}

type StringValueStruct struct {
	StringValue string `json:"stringValue"`
}

type TimestampValueStruct struct {
	TimestampValue time.Time `json:"timestampValue"`
}

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

type FirestorePersonEvent struct {
	OldValue   FirestorePersonValue `json:"oldValue"`
	Value      FirestorePersonValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

type FirestorePersonValue struct {
	CreateTime time.Time             `json:"createTime"`
	Fields     PersonFirestoreFields `json:"fields"`
	Name       string                `json:"name"`
	UpdateTime time.Time             `json:"updateTime"`
}
