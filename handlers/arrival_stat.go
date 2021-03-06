package handlers

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firesearch"
	"context"
	"fmt"
	"strings"
)

// ArrivalStatEventResult is the output of the handler that reacts to arrival events
type ArrivalStatEventResult struct {
	Event       models.FirestoreArrivalEvent
	ArrivalStat models.ArrivalStat
}

// ArrivalStatEvent handles events from the arrivals collection
func ArrivalStatEvent(ctx context.Context, event models.FirestoreArrivalEvent, store firesearch.ArrivalsStore) (ArrivalStatEventResult, error) {
	fields := event.Value.Fields
	year, month, _ := fields.DateOfArrival.TimestampValue.Date()
	date := fields.DateOfArrival.TimestampValue.Format("2006-01-02")
	d := strings.Split(date, "-")
	arrivalStat := models.ArrivalStat{
		ID:                   fields.ID.StringValue,
		Date:                 strings.Join(d, ""),
		Year:                 year,
		Month:                fmt.Sprintf("%d-%d", year, month),
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
