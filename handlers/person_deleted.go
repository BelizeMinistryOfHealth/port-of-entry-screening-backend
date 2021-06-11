package handlers

import (
	"bz.moh.epi/poebackend/repository/firesearch"
	"bz.moh.epi/poebackend/repository/firestore"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// PersonDeletedArgs are the arguments for deleting a person
type PersonDeletedArgs struct {
	PersonFiresearchStore firesearch.PersonStore
	ArrivalStoreService   *firestore.ArrivalsStoreService
	AddressStoreService   *firestore.AddressStoreService
}

// PersonDeleted is a handler that gets called when a person is deleted from
// the persons collection.
func PersonDeleted(ctx context.Context, args PersonDeletedArgs, ID string) error {
	personStore := args.PersonFiresearchStore
	if err := personStore.DeletePerson(ctx, ID); err != nil {
		return fmt.Errorf("PersonDeleted handler failed: %w", err)
	}
	if err := args.ArrivalStoreService.DeleteArrival(ctx, ID); err != nil {
		log.WithFields(log.Fields{"personID": ID}).WithError(err)
	}
	if err := args.AddressStoreService.DeleteAddress(ctx, ID); err != nil {
		log.WithFields(log.Fields{"personID": ID}).WithError(err)
	}
	return nil
}
