package firesearch

import (
	"bz.moh.epi/poebackend/models"
	"context"
	"fmt"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
)

// PersonStore encapsulates a Firesearch connection and actions that can be
// performed on the person index.
type PersonStore struct {
	Service Service
}

// CreatePerson creates a new person record in Firesearch.
func (c *PersonStore) CreatePerson(ctx context.Context, person models.PersonalInfo) error {
	putDocReq := firesearch.PutDocRequest{
		IndexPath: c.Service.IndexPath,
		Doc: firesearch.Doc{
			ID: person.ID,
			SearchFields: []firesearch.SearchField{
				{
					Key:   "firstName",
					Value: person.FirstName,
					Store: true,
				},
				{
					Key:   "lastName",
					Value: person.LastName,
					Store: true,
				},
				{
					Key:   "fullName",
					Value: person.FullName,
					Store: true,
				},
			},
			Fields: []firesearch.Field{
				{
					Key:   "ID",
					Value: person.ID,
				},
				{
					Key:   "middleName",
					Value: person.MiddleName,
				},
				{
					Key:   "dob",
					Value: person.Dob,
				},
				{
					Key:   "gender",
					Value: person.Gender,
				},
				{
					Key:   "nationality",
					Value: person.Nationality,
				},
				{
					Key:   "phoneNumbers",
					Value: person.PhoneNumbers,
				},
				{
					Key:   "passportNumber",
					Value: person.PassportNumber,
				},
				{
					Key:   "occupation",
					Value: person.Occupation,
				},
				{
					Key:   "email",
					Value: person.Email,
				},
			},
		},
	}

	_, err := c.Service.IndexService.PutDoc(ctx, putDocReq)
	if err != nil {
		return fmt.Errorf("CreatePerson failed: %w", err)
	}
	return nil
}

// DeletePerson deletes a person's search index.
// It is triggered when a person is deleted from the `persons` collection.
func (c *PersonStore) DeletePerson(ctx context.Context, ID string) error {
	deleteDocReq := firesearch.DeleteDocRequest{
		IndexPath: c.Service.IndexPath,
		ID:        ID,
	}
	_, err := c.Service.IndexService.DeleteDoc(ctx, deleteDocReq)
	if err != nil {
		return fmt.Errorf("DeletePerson() failed: %w", err)
	}
	return nil
}
