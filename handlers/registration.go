package handlers

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firestore"
	"context"
	"encoding/json"
	"fmt"
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
	Companions   []models.PersonalInfo  `json:"companions"`
}

// RegistrationHandler creates a new registration
func RegistrationHandler(args RegistrationArgs, w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithFields(log.Fields{
			"request": r.Body,
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

	companionErr := saveCompanions(r.Context(), args, req)
	if companionErr != nil {
		log.WithFields(log.Fields{
			"request": req,
		}).WithError(companionErr).Error("RegistrationHandler(): creating companions failed")
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

func saveCompanions(ctx context.Context, args RegistrationArgs, req RegistrationRequest) error {
	companions := req.Companions
	for _, c := range companions {
		editor := models.Editor{
			ID:    "00000",
			Email: "system@openstep.net",
		}
		now := time.Now()
		// Create person
		c.CreatedBy = editor
		c.ModifiedBy = editor
		c.Created = now
		c.Modified = now
		if err := args.PersonStoreService.CreatePerson(ctx, req.PersonalInfo); err != nil {
			return fmt.Errorf("failed to create companion record (%s): %w", c.ID, err)
		}

		// Create address
		req.Address.ID = c.ID
		if err := args.AddressStoreService.CreateAddress(ctx, req.Address); err != nil {
			return fmt.Errorf("failed to create address for companion (%s): %w", c.ID, err)
		}

		req.ArrivalInfo.ID = c.ID
		if err := args.ArrivalStoreService.CreateArrival(ctx, req.ArrivalInfo); err != nil {
			return fmt.Errorf("failed to create arrival info for companion (%s): %w", c.ID, err)
		}
	}

	return nil
}
