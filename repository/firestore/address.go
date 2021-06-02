package firestore

import (
	"bz.moh.epi/poebackend/models"
	fs "cloud.google.com/go/firestore"
	"context"
	"fmt"
)

// AddressStoreService is service for interacting with persisted address data
type AddressStoreService struct {
	db         *DB
	collection string
	colRef     *fs.CollectionRef
}

// CreateAddressStoreService constructor
func CreateAddressStoreService(db *DB, collection string) *AddressStoreService {
	return &AddressStoreService{
		db:         db,
		collection: collection,
		colRef:     db.Client.Collection(collection),
	}
}

// GetByID retrieves address by its id
func (p *AddressStoreService) GetByID(ctx context.Context, ID string) (models.AddressInBelize, error) {
	dsnap, err := p.colRef.Doc(ID).Get(ctx)
	if err != nil {
		return models.AddressInBelize{}, fmt.Errorf("AddressStoreService.GetByID failed: %w", err)
	}

	if dsnap == nil || !dsnap.Exists() {
		return models.AddressInBelize{}, ErrNoResult
	}

	var address models.AddressInBelize
	decodeErr := dsnap.DataTo(&address)
	if decodeErr != nil {
		return models.AddressInBelize{}, fmt.Errorf("AddressStoreService.GetByID decoding error: %w", err)
	}
	return address, nil
}
