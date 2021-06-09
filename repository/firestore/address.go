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

func (s *AddressStoreService) CreateAddress(ctx context.Context, address models.AddressInBelize) error {
	_, err := s.colRef.Doc(address.ID).Set(ctx, address)
	if err != nil {
		return fmt.Errorf("AddressStoreService.CreateAddress: failed: %w", err)
	}
	return nil
}

func (s *AddressStoreService) DeleteAddress(ctx context.Context, ID string) error {
	_, err := s.colRef.Doc(ID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("DeleteAddress failed: %w", err)
	}
	return nil
}
