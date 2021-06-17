package firesearch

import (
	"bz.moh.epi/poebackend/models"
	"context"
	"fmt"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
	"time"
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
					Key:   "fullName",
					Value: person.FullName,
					Store: true,
				},
				{
					Key:   "passportNumber",
					Value: person.PassportNumber,
					Store: true,
				},
				{
					Key:   "email",
					Value: person.Email,
					Store: true,
				},
				{
					Key:   "otherTravelDocumentId",
					Value: person.OtherTravelDocumentID,
					Store: true,
				},
				{
					Key:   "year",
					Value: fmt.Sprintf("%d", person.Created.Year()),
					Store: true,
				},
				{
					Key:   "month",
					Value: fmt.Sprintf("%d", person.Created.Month()),
					Store: true,
				},
			},
			Fields: []firesearch.Field{
				{
					Key:   "id",
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
					Key:   "occupation",
					Value: person.Occupation,
				},
				{
					Key:   "created",
					Value: person.Created,
				},
				{
					Key:   "day",
					Value: person.Created.Day(),
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

// SearchByName searches for a person record by their name
func (c *PersonStore) SearchByName(ctx context.Context, accessKey, portOfEntry, name string) (
	[]PersonSearchResult, error) {
	filters := []firesearch.Field{
		{
			Key:   "portOfEntry",
			Value: portOfEntry,
		},
	}

	searchReq := firesearch.SearchRequest{
		Query: firesearch.SearchQuery{
			IndexPath: c.Service.IndexPath,
			AccessKey: accessKey,
			Limit:     50,
			Text:      name,
			Filters:   filters,
			Select: []string{
				"id",
				"firstName",
				"lastName",
				"middleName",
				"fullName",
				"gender",
				"dob",
				"nationality",
				"occupation",
				"passportNumber",
				"email",
			},
			SearchFields: []string{},
		},
	}
	searchResp, err := c.Service.IndexService.Search(ctx, searchReq)
	if err != nil {
		return []PersonSearchResult{}, fmt.Errorf("SearchByName() failed: %w", err)
	}
	hits := searchResp.Hits
	var persons []PersonSearchResult
	for _, h := range hits { //nolint:typecheck
		fields := h.Fields
		persons = append(persons, PersonSearchResult{
			ID:             GetField(fields, "id").(string),
			FirstName:      GetField(fields, "firstName").(string),
			MiddleName:     GetField(fields, "middleName").(string),
			LastName:       GetField(fields, "lastName").(string),
			FullName:       GetField(fields, "fullName").(string),
			Gender:         GetField(fields, "gender").(string),
			Nationality:    GetField(fields, "nationality").(string),
			Dob:            GetField(fields, "dob").(time.Time),
			Occupation:     GetField(fields, "occupation").(string),
			PassportNumber: GetField(fields, "passportNumber").(string),
			Email:          GetField(fields, "email").(string),
		})

	}
	return persons, nil
}

// PersonSearchResult is the result from searching for persons in firesearch
type PersonSearchResult struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	MiddleName     string    `json:"middleName"`
	FullName       string    `json:"fullName"`
	Gender         string    `json:"gender"`
	Nationality    string    `json:"nationality"`
	Dob            time.Time `json:"dob"`
	Occupation     string    `json:"occupation"`
	PassportNumber string    `json:"passportNumber"`
	Email          string    `json:"email"`
}
