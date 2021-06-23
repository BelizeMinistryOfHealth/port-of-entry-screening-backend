package handlers

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firesearch"
	"context"
	"fmt"
)

// ArrivalStatEventResult is the output of the handler that reacts to arrival events
type ArrivalStatEventResult struct {
	Event       models.FirestoreArrivalEvent
	ArrivalStat models.ArrivalStat
}

// ArrivalStatEvent handles events from the arrivals collection
func ArrivalStatEvent(ctx context.Context, event models.FirestoreArrivalEvent, store firesearch.ArrivalsStore) (ArrivalStatEventResult, error) {
	fields := event.Value.Fields
	month := fields.DateOfArrival.TimestampValue.Month()
	year := fields.DateOfArrival.TimestampValue.Year()
	arrivalStat := models.ArrivalStat{
		ID:                   fields.ID.StringValue,
		Date:                 fields.DateOfArrival.TimestampValue.Format("2006-01-02"),
		Year:                 year,
		Month:                fmt.Sprintf("%d-%s", year, month.String()),
		PortOfEntry:          fields.PortOfEntry.StringValue,
		CountryOfEmbarkation: fields.CountryOfEmbarkation.StringValue,
		PurposeOfTrip:        fields.PurposeOfTrip.StringValue,
	}
	err := store.PutDoc(ctx, arrivalStat)
	if err != nil {
		return ArrivalStatEventResult{Event: event}, fmt.Errorf("failed to save arrival stat: %w", err)
	}
	return ArrivalStatEventResult{
		Event:       event,
		ArrivalStat: arrivalStat,
	}, nil
}

func ArrivalDeleted(ctx context.Context, arrivalStore firesearch.ArrivalsStore, ID string) error {
	if err := arrivalStore.DeleteDoc(ctx, ID); err != nil {
		return fmt.Errorf("ArrivalDeleted() failed: %w", err)
	}
	return nil
}
