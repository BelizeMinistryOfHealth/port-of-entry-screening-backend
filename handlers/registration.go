package handlers

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firestore"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// RegistrationArgs are the arguments for the Registration Handler
type RegistrationArgs struct {
	PersonStoreService  *firestore.PersonStoreService
	ArrivalStoreService *firestore.ArrivalsStoreService
	AddressStoreService *firestore.AddressStoreService
}

// RegistrationRequest is the JSON request posted for registration
type RegistrationRequest struct {
	PersonalInfo models.PersonalInfo    `json:"personalInfo"`
	ArrivalInfo  models.ArrivalInfo     `json:"arrivalInfo"`
	Address      models.AddressInBelize `json:"address"`
}

// RegistrationHandler creates a new registration
func RegistrationHandler(args RegistrationArgs, w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithFields(log.Fields{
			"request": req,
		}).WithError(err).Error("RegistrationHandler(): decoding failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	editor := models.Editor{
		ID:    "00000",
		Email: "system@openstep.net",
	}
	now := time.Now()
	// Create person
	req.PersonalInfo.CreatedBy = editor
	req.PersonalInfo.ModifiedBy = editor
	req.PersonalInfo.Created = now
	req.PersonalInfo.Modified = now
	if err := args.PersonStoreService.CreatePerson(r.Context(), req.PersonalInfo); err != nil {
		log.WithFields(log.Fields{
			"request": req,
		}).WithError(err).Error("RegistrationHandler(): creating person failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Create address
	req.Address.CreatedBy = editor
	req.Address.ModifiedBy = editor
	req.Address.Created = &now
	req.Address.Modified = &now
	if err := args.AddressStoreService.CreateAddress(r.Context(), req.Address); err != nil {
		log.WithFields(log.Fields{
			"request": req,
		}).WithError(err).Error("RegistrationHandler(): creating address failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Create arrival
	req.ArrivalInfo.CreatedBy = editor
	req.ArrivalInfo.ModifiedBy = editor
	req.ArrivalInfo.Created = &now
	req.ArrivalInfo.Modified = &now
	if err := args.ArrivalStoreService.CreateArrival(r.Context(), req.ArrivalInfo); err != nil {
		log.WithFields(log.Fields{
			"request": req,
		}).WithError(err).Error("RegistrationHandler(): creating arrival failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode("OK")
	if err != nil {
		log.WithFields(log.Fields{
			"req": req,
		}).WithError(err).Error("encoding response failed")
		return
	}
}
