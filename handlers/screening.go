package handlers

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firestore"
	"bz.moh.epi/poebackend/repository/godata"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func ScreeningEventHandler(
	ctx context.Context,
	event models.FirestoreScreeningEvent,
	personStore *firestore.PersonStoreService,
	arrivalStore *firestore.ArrivalsStoreService,
	addressStore *firestore.AddressStoreService) error {

	godataURL := os.Getenv("GODATA_URL")
	godataUsername := os.Getenv("GODATA_USERNAME")
	godataPassword := os.Getenv("GODATA_PASSWORD")
	outbreakID := os.Getenv("OUTBREAK_ID")

	screening := event.Value.Fields.ToScreening()
	arrivalInfo, arrivalErr := arrivalStore.GetByID(ctx, screening.ID)
	if arrivalErr != nil {
		return fmt.Errorf("error: can not push to GoData without arrival info: %w", arrivalErr)
	}
	address, addressErr := addressStore.GetByID(ctx, screening.ID)
	if addressErr != nil {
		return fmt.Errorf("error: can not push to GoData without address info: %w", addressErr)
	}

	personalInfo, personErr := personStore.GetByID(ctx, screening.ID)
	if personErr != nil {
		return fmt.Errorf("error: can not push to GoData without personal info: %w", personErr)
	}
	// Retrieve Token
	godataToken, tokenErr := godata.GetGodataToken(godataUsername, godataPassword, godataURL)
	if tokenErr != nil {
		return fmt.Errorf("failed to retrieve godata token: %w", tokenErr)
	}
	caseID, err := godata.GetCaseByVisualId(screening.ID, godata.Options{
		Url:   godataURL,
		Token: godataToken,
	})
	arg := godata.GodataCaseArg{
		PersonalInfo: personalInfo,
		Screening:    screening,
		ArrivalInfo:  arrivalInfo,
		Address:      address,
		VisualId:     personalInfo.ID,
	}
	if err != nil {
		// PUT
		if err := godata.UpdateGoDataCase(arg, caseID.ID, godata.Options{
			Url:        godataURL,
			Token:      godataToken,
			OutbreakId: outbreakID,
		}); err != nil {
			log.WithFields(log.Fields{
				"screening":    screening,
				"personalInfo": personalInfo,
				"caseID":       caseID,
			}).WithError(err).Error("failed to update godata case")
			return fmt.Errorf("error updating GoData case: %w", err)
		}
	}

	if err := godata.PushToGoData(arg,
		godata.Options{Url: godataURL, Token: godataToken, OutbreakId: outbreakID}); err != nil {
		log.WithFields(log.Fields{
			"screening":    screening,
			"personalInfo": personalInfo,
			"caseID":       caseID,
		}).WithError(err).Error("failed to update godata case")
		return fmt.Errorf("error creating new GoData case: %w", err)
	}
	// Check if POST or PUT
	// Do upload
	return nil
}
