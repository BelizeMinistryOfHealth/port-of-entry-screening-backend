package firestore

import fs "cloud.google.com/go/firestore"

// VaccinationService represents a service for persisting vaccination information
type VaccinationService struct {
	db         *DB
	collection string
	colRef     *fs.CollectionRef
}

// CreateVaccinationService creates a new service
func CreateVaccinationService(db *DB, collection string) *VaccinationService {
	return &VaccinationService{
		db:         db,
		collection: "vaccines",
		colRef:     db.Client.Collection(collection),
	}
}
