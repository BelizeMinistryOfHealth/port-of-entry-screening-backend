package handlers

import (
	"bz.moh.epi/poebackend/repository/firesearch"
	"context"
	"fmt"
)

// PersonDeleted is a handler that gets called when a person is deleted from
// the persons collection.
func PersonDeleted(ctx context.Context, personStore firesearch.PersonStore, ID string) error {
	if err := personStore.DeletePerson(ctx, ID); err != nil {
		return fmt.Errorf("PersonDeleted handler failed: %w", err)
	}
	return nil
}
