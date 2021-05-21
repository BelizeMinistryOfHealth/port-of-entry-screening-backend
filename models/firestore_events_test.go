package models

import (
	"context"
	"github.com/cloudevents/sdk-go/v2/event/datacodec/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const event1 = `{"fields": {
                "gender": {
                    "stringValue": "Male"
                },
                "email": {
                    "stringValue": ""
                },
                "bhisNumber": {
                    "stringValue": ""
                },
                "nationality": {
                    "stringValue": "BZ"
                },
                "middleName": {
                    "stringValue": "Midd"
                },
                "otherTravelDocument": {
                    "stringValue": "some"
                },
                "createdBy": {
                    "mapValue": {
                        "fields": {
                            "email": {
                                "stringValue": "uris77@gmail.com"
                            },
                            "id": {
                                "stringValue": "zSTdPRfJUCaQVwa5vGNN0273hfl2"
                            }
                        }
                    }
                },
                "modified": {
                    "timestampValue": "2021-05-19T13:40:27.868Z"
                },
                "otherTravelDocumentId": {
                    "stringValue": "asdf9asfasdf"
                },
                "passportNumber": {
                    "stringValue": "768888888"
                },
                "lastName": {
                    "stringValue": "Guerra"
                },
                "phoneNumbers": {
                    "stringValue": "67123123"
                },
                "modifiedBy": {
                    "nullValue": null
                },
                "created": {
                    "timestampValue": "2021-05-19T13:40:27.868Z"
                },
                "firstName": {
                    "stringValue": "Uris"
                },
                "dob": {
                    "stringValue": "2008-02-07"
                },
                "fullName": {
                    "stringValue": "Uris Midd Guerra"
                }
            }
}`

func Test_MarshalPersonFields(t *testing.T) {
	ctx := context.Background()
	var fields PersonFirestoreFields
	if err := json.Decode(ctx, []byte(event1), &fields); err != nil {
		t.Fatalf("failed to marshal person event: %v", err)
	}
	createdBy := fields.Fields.CreatedBy.MapValueStruct.Fields
	if len(createdBy.ID.StringValue) == 0 {
		t.Errorf("createdBy.ID should not be empty")
	}
	if createdBy.Email.StringValue != "uris77@gmail.com" {
		t.Errorf("want: uris77@gmail.com, got: %s", createdBy.Email.StringValue)
	}
	//createdDate :=
	//t.Logf("created: %v", fields.Fields.Created.TimestampValue.Year())
	assert.Equal(t, "2021-05-19", fields.Fields.Created.TimestampValue.Format("2006-01-02"))
}
