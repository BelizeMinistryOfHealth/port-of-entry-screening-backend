package persistence

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"bz.epi.covid.screen/arrivals/domain"
)

type FirestoreClient struct {
	client    *firestore.Client
	ctx       context.Context
	projectID string
}

type FirestoreDb interface {
	UpsertArrival(i domain.Arrival) error
	CreateClient(ctx context.Context, projectID string) FirestoreClient
	FindByName(first string, last string) ([]domain.Arrival, error)
	FindByPortOfEntry(col, loc string) ([]domain.Arrival, error)
	Close() error
}

func CreateClient(ctx context.Context, projectID string) *FirestoreClient {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &FirestoreClient{
		client:    client,
		ctx:       ctx,
		projectID: projectID,
	}
}

func (c FirestoreClient) Close() error {
	return c.client.Close()
}

func (c FirestoreClient) UpsertArrival(col string, arrivals []domain.Arrival) error {
	batch := c.client.Batch()
	for _, v := range arrivals {
		v.AddressInBelize.District = strings.Title(strings.ToLower(v.AddressInBelize.District))
		v.AddressInBelize.Municipality = strings.Title(strings.ToLower(v.AddressInBelize.Municipality))
		ref := c.client.Collection(col).Doc(v.Id)
		batch.Set(ref, v)
	}
	//_, err := c.client.Collection(col).Doc(arrivals.PersonalInfo.PassportNumber).Set(c.ctx, arrivals)
	_, err := batch.Commit(c.ctx)
	return err
}

func (c FirestoreClient) FindByName(col string, first string, last string) ([]domain.Arrival, error) {
	arrivals := []domain.Arrival{}
	it := c.client.Collection(col).
		Where("PersonalInfo.FirstName", "==", first).
		Where("PersonalInfo.LastName", "==", last).
		Documents(c.ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var a domain.Arrival
		doc.DataTo(&a)

		arrivals = append(arrivals, a)
	}

	return arrivals, nil
}

func (c FirestoreClient) FindByPortOfEntry(col, loc string) ([]domain.Arrival, error) {
	arrivals := []domain.Arrival{}
	it := c.client.Collection(col).
		Where("ArrivalInfo.PortOfEntry", "==", loc).
		Documents(c.ctx)

	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var a domain.Arrival
		doc.DataTo(&a)
		arrivals = append(arrivals, a)
	}
	return arrivals, nil
}
