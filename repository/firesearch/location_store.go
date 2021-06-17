package firesearch

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"strings"
)

// District is a string representation of districts in Belize
type District string

const (
	// Belize district
	Belize District = "Belize"
	// Corozal District
	Corozal District = "Corozal"
	// OrangeWalk District
	OrangeWalk District = "Orange Walk"
	// Cayo District
	Cayo District = "Cayo"
	// StannCreek District
	StannCreek District = "Stann Creek"
	// Toledo District
	Toledo District = "Toledo"
)

// Community is a town, village or city in a District
type Community struct {
	ID       string   `json:"id" firestore:"id"`
	Name     string   `json:"name" firestore:"name"`
	District District `json:"district" firestore:"district"`
}

// LocationStore represents a location's database connection
type LocationStore struct {
	db         *firestore.Client
	collection string
}

// CreateStore initiates a firebase connection. This store can then be used to manipulate
// the location information inside Firebase.
func CreateStore(ctx context.Context, projectID string, coll string) (*LocationStore, error) {
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failure creating new firebae client: %w", err)
	}

	store := &LocationStore{
		db:         c,
		collection: coll,
	}
	return store, nil
}

// SaveCommunity saves a community to the database.
func (s *LocationStore) SaveCommunity(ctx context.Context, com Community) error {
	_, err := s.db.Collection(s.collection).Doc(com.ID).Set(ctx, com)
	if err != nil {
		return fmt.Errorf("LocationStore.SaveCommunity failed: %w", err)
	}
	return nil
}

// FindByDistrict finds all the communities that are found in a particular district.
func (s *LocationStore) FindByDistrict(ctx context.Context, district District) ([]Community, error) {
	var comms []Community
	coll := s.collection
	it := s.db.Collection(coll).
		Where("district", "==", district).
		OrderBy("name", firestore.Desc).
		Documents(ctx)
	for {
		doc, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error finding communities by District (%v): %w", district, err)
		}
		var community Community
		if err := doc.DataTo(&community); err != nil {
			return nil, fmt.Errorf("LocationStore.FindByDistrict failed while decoding data: %w", err)
		}
		comms = append(comms, community)
	}
	return comms, nil
}

// FindByID finds a community for the matching id.
func (s *LocationStore) FindByID(ctx context.Context, id string) (Community, error) {
	r, err := s.db.Collection(s.collection).Doc(id).Get(ctx)
	if err != nil {
		return Community{}, fmt.Errorf("error retrieving from firebase: %w", err)
	}
	var c Community
	err = r.DataTo(c)
	if err != nil {
		return Community{}, fmt.Errorf("error marshalling community from firebase: %w", err)
	}
	return c, nil
}

// FindByName finds a community that matches the specified name in a particular district.
func (s *LocationStore) FindByName(ctx context.Context, name, district string) (Community, error) {
	var n = name
	if name == "Santa Elena - Cayo2" {
		n = "Santa Elena"
	}
	it := s.db.Collection(s.collection).
		Where("name", "==", strings.Trim(n, "")).
		Where("district", "==", district).
		Limit(1).
		Documents(ctx)
	var c Community

	for {
		doc, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return Community{}, fmt.Errorf("error retrieving from firebase: %w", err)
		}
		if err := doc.DataTo(&c); err != nil {
			return c, fmt.Errorf("LocationStore.FindByName: error decoding community %w", err)
		}
	}
	return c, nil
}
