package poebackend

import (
	"bz.moh.epi/poebackend/models"
	"context"
	"time"
)

// PersonStoreService represents a service for managing persons entering through a port of entry.
type PersonStoreService interface {
	// CreatePerson creates a new person.
	CreatePerson(ctx context.Context, person models.Person) error

	// UpdatePerson updates the person's arrival and personal information.
	// Note that screening and vaccination information is updated separately.
	UpdatePerson(ctx context.Context, person models.Person) error

	// GetByID retrieves a person that matches the provided ID.
	// If no record is found, an error is returned.
	GetByID(ctx context.Context, id string) (models.Person, error)
}

// SearchService represents a service that performs full text search.
type SearchService interface {
	// SearchPerson conducts a full text search for a person. The txt could match the full name,
	// or any part of the name, as well as their QR code, etc.
	SearchPerson(ctx context.Context, txt string) ([]models.Person, error)

	// SearchByDate lists persons who will arrive on the given date
	SearchByDate(ctx context.Context, date time.Time) ([]models.Person, error)
}
