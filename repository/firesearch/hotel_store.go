package firesearch

import (
	"bz.moh.epi/poebackend/models"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

// HotelStore represents the hotel's database connection
type HotelStore struct {
	db         *firestore.Client
	collection string
}

// CreateHotelStore initiates a firebase connection. This store can then be used to manipulate
// the hotel information inside Firebase.
func CreateHotelStore(ctx context.Context, projectID string, coll string) (*HotelStore, error) {
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failure creating new firebase client: %w", err)
	}

	store := &HotelStore{
		db:         c,
		collection: coll,
	}
	return store, nil
}

// FindHotelByID retrieves a hotel with a specified ID
func (c *HotelStore) FindHotelByID(ctx context.Context, ID string) (models.AddressInBelize, error) {
	dsnap, err := c.db.Collection(c.collection).Doc(ID).Get(ctx)
	if err != nil {
		return models.AddressInBelize{}, fmt.Errorf("failure retrieving hotel by ID: %w", err)
	}
	var address models.AddressInBelize
	if err := dsnap.DataTo(&address); err != nil {
		return models.AddressInBelize{}, fmt.Errorf("HotelStore.FindHotelByID failed decoding data: %w", err)
	}
	return address, nil
}
