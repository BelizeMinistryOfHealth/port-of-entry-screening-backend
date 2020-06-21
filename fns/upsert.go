package fns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"bz.epi.covid.screen/arrivals/domain"
	"bz.epi.covid.screen/arrivals/persistence"
)

func UpsertArrival(dbClient *persistence.FirestoreClient, collection string, r *http.Request) ([]domain.Arrival, error) {
	var newArrivals []domain.ArrivalRequest

	// Read body
	body := r.Body
	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not parse the body posted: %v", err)
	}

	logrus.WithFields(logrus.Fields{
		"body": string(b),
	}).Info("Got Upsert request for arrivals")
	err = json.Unmarshal(b, &newArrivals)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"body":  string(b),
			"error": err,
		}).Error("could not unmarshall the body posted")
		return nil, fmt.Errorf("could not unmarshall the body posted: %v", err)
	}

	arrivals := domain.HydrateCompanions(FilterSyncedArrivals(newArrivals))

	err = dbClient.UpsertArrival(collection, arrivals)
	if err != nil {
		return nil, fmt.Errorf("error saving newArrivals information: %v", err)
	}

	return arrivals, nil
}

// FilterSyncedArrivals removes any arrival that has already been synced. It makes sure
// we do not over write data with stale data.
func FilterSyncedArrivals(arrivals []domain.ArrivalRequest) []domain.ArrivalRequest {
	var newArrivals []domain.ArrivalRequest

	for _, v := range arrivals {
		if v.SyncStatus == domain.NEW_ARRIVAL {
			newArrivals = append(newArrivals, v)
		}
	}

	return newArrivals
}
