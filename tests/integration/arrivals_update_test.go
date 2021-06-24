package integration

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firestore"
	"context"
	"os"
	"testing"
	"time"
)

func TestArrivalsUpdate(t *testing.T) {
	ctx := context.Background()
	// connect to firestore
	projectID := os.Getenv("PROJECT_ID")
	firestoreDb, err := firestore.CreateFirestoreDB(ctx, projectID)

	if err != nil {
		t.Fatalf("Failed to create firestore db: %v", err)
	}

	arrivalStoreService := firestore.CreateArrivalsStoreService(firestoreDb, "arrivals")
	//now := time.Now().Format("2006-01-02")
	date, err := time.Parse("2006-01-02", "2021-06-20")
	if err != nil {
		t.Fatalf("error parsing date: %v", err)
	}

	arrivals, err := arrivalStoreService.FindByDateOfArrival(ctx, date)
	if err != nil {
		t.Fatalf("error retrieving arrivals: %v", err)
	}

	var touchedArrivals []models.ArrivalInfo
	for _, arrival := range arrivals {
		t.Logf("touching from %v to %v", arrival.Touch, !arrival.Touch)
		arrival.Touch = !arrival.Touch
		touchedArrivals = append(touchedArrivals, arrival)
	}

	batchErr := arrivalStoreService.BatchUpdate(ctx, arrivals)
	if batchErr != nil {
		t.Errorf("failed to update: %v", batchErr)
	}

	t.Logf("\n\ntouched: %v", touchedArrivals)
	t.Logf("\n\narrivals: %v", arrivals)

}
