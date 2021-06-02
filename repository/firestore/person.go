package firestore

import (
	"bz.moh.epi/poebackend/models"
	fs "cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
)

// PersonStoreService represents a service for persisting persons
type PersonStoreService struct {
	db         *DB
	collection string
	colRef     *fs.CollectionRef
}

// CreatePersonService instantiates a new person service
func CreatePersonService(db *DB, collection string) *PersonStoreService {
	return &PersonStoreService{
		db:         db,
		collection: collection,
		colRef:     db.Client.Collection(collection),
	}
}

// CreatePerson inserts a new person to the database
func (p *PersonStoreService) CreatePerson(ctx context.Context, person models.Person) error {
	_, err := p.colRef.Doc(person.ID).Set(ctx, person)
	if err != nil {
		return fmt.Errorf("PersonStoreService.CreatePerson: failed: %w", err)
	}
	return nil
}

// UpdatePerson will update a person in firestore.
func (p *PersonStoreService) UpdatePerson(ctx context.Context, person models.Person) error {
	ref := p.colRef.Doc(person.ID)
	_, err := ref.Update(ctx, []fs.Update{
		{
			Path:  "personalInfo",
			Value: person.PersonalInfo,
		},
		{
			Path:  "arrival.arrivalInfo",
			Value: person.Arrival.ArrivalInfo,
		},
		{
			Path:  "arrival.hotelAddress",
			Value: person.Arrival.HotelAddress,
		},
		{
			Path:  "arrival.address",
			Value: person.Arrival.Address,
		},
		{
			Path:  "arrival.purposeOfTrip",
			Value: person.Arrival.PurposeOfTrip,
		},
	})

	if err != nil {
		return fmt.Errorf("PersonStoreService.UpdatePerson: failed to update person %w", err)
	}
	return nil
}

// ErrNoResult is an error when no record is found
var ErrNoResult = errors.New("no record found")

// GetByID retrieves a person by its ID
func (p *PersonStoreService) GetByID(ctx context.Context, id string) (models.Person, error) {
	dsnap, err := p.colRef.Doc(id).Get(ctx)
	if !dsnap.Exists() {
		return models.Person{}, ErrNoResult
	}
	if err != nil {
		return models.Person{}, fmt.Errorf("PersonStoreService.GetByID: error retrieving person: %w", err)
	}
	var person models.Person
	decodeErr := dsnap.DataTo(&person)
	if decodeErr != nil {
		return models.Person{}, fmt.Errorf("PersonStoreService.GetByID: error decoding person record: %w", decodeErr)
	}
	return person, nil
}
