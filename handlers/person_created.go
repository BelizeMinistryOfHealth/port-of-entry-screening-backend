package handlers

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firesearch"
	"context"
	"fmt"
	"github.com/bearbin/go-age"
	"time"
)

// PersonCreatedResult is the output of PersonCreated
type PersonCreatedResult struct {
	Age   int
	Event models.FirestorePersonEvent
}

// PersonCreated is a handler that gets triggered when a new person is created.
// It will create a record in Firesearch for this person.
func PersonCreated(ctx context.Context, event models.FirestorePersonEvent, personStore firesearch.PersonStore) (PersonCreatedResult, error) {
	dob := event.Value.Fields.Dob.StringValue
	dobDate, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return PersonCreatedResult{}, fmt.Errorf("could not parse dob: %w", err)
	}

	if err := personStore.CreatePerson(ctx, event.Value.Fields.ToPerson()); err != nil {
		return PersonCreatedResult{}, fmt.Errorf("PersonCreated() failed: %w", err)
	}
	return PersonCreatedResult{
		Age:   age.AgeAt(dobDate, event.Value.Fields.Created.TimestampValue),
		Event: event,
	}, nil
}
