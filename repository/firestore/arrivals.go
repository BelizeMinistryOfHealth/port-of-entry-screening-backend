package firestore

import (
	"bz.moh.epi/poebackend/models"
	fs "cloud.google.com/go/firestore"
	"context"
	"fmt"
)

// ArrivalsStoreService store for interacting with persisted arrivals data
type ArrivalsStoreService struct {
	db         *DB
	collection string
	colRef     *fs.CollectionRef
}

// CreateArrivalsStoreService instantiates a new arrival service
func CreateArrivalsStoreService(db *DB, collection string) *ArrivalsStoreService {
	return &ArrivalsStoreService{
		db:         db,
		collection: collection,
		colRef:     db.Client.Collection(collection),
	}
}

// GetByID retrieves an arrival by its ID
func (p *ArrivalsStoreService) GetByID(ctx context.Context, ID string) (models.ArrivalInfo, error) {
	dsnap, err := p.colRef.Doc(ID).Get(ctx)
	if err != nil {
		return models.ArrivalInfo{}, fmt.Errorf("ArrivalsStoreService.GetByID: error retrieving arrival: %w", err)
	}
	if dsnap == nil || !dsnap.Exists() {
		return models.ArrivalInfo{}, ErrNoResult
	}
	var arrival models.ArrivalInfo
	decodeErr := dsnap.DataTo(&arrival)
	if decodeErr != nil {
		return models.ArrivalInfo{}, fmt.Errorf("ArrivalsStoreService.GetByID: error decoding arrival: %w", err)
	}
	return arrival, nil
}

// CreateArrival persists an arrival in Firestore
func (p *ArrivalsStoreService) CreateArrival(ctx context.Context, arrival models.ArrivalInfo) error {
	_, err := p.colRef.Doc(arrival.ID).Set(ctx, arrival)
	if err != nil {
		return fmt.Errorf("ArrivalStoreService.CreateArrival: failed: %w", err)
	}
	return nil
}

// DeleteArrival deletes an arrival from Firestore
func (p *ArrivalsStoreService) DeleteArrival(ctx context.Context, ID string) error {
	_, err := p.colRef.Doc(ID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("DeleteArrival failed: %w", err)
	}
	return nil
}
