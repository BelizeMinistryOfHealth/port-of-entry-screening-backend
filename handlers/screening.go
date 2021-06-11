package handlers

import (
	"bz.moh.epi/poebackend/models"
	"bz.moh.epi/poebackend/repository/firestore"
	"bz.moh.epi/poebackend/repository/godata"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

func hyphenateString(s string) string {
	split := strings.Split(s, " ")
	return strings.Join(split, "-")
}

// ScreeningEventHandler is a handler that gets triggered when a screening record is created or modified.
// It updates or creates a new GoData Record.
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
	ID := fmt.Sprintf("%s#%s", hyphenateString(arrivalInfo.PortOfEntry), personalInfo.ID)
	httpClient := http.Client{}
	goAPI := godata.NewAPI(godataURL, &httpClient)
	caseID, err := goAPI.GetCaseByVisualID(ID, godata.Options{
		URL:   godataURL,
		Token: godataToken,
	})

	log.WithFields(log.Fields{
		"ID":     ID,
		"CaseID": caseID,
	}).Info("Retrieved godata case")

	arg := godata.CaseArg{
		PersonalInfo: personalInfo,
		Screening:    screening,
		ArrivalInfo:  arrivalInfo,
		Address:      address,
		VisualID:     ID,
	}
	if err != nil || len(caseID.ID) > 0 {
		// PUT
		if err := goAPI.UpdateCase(arg, caseID.ID, godata.Options{
			URL:        godataURL,
			Token:      godataToken,
			OutbreakID: outbreakID,
		}); err != nil {
			log.WithFields(log.Fields{
				"screening":    screening,
				"personalInfo": personalInfo,
				"caseID":       caseID,
			}).WithError(err).Error("failed to update godata case")
			return fmt.Errorf("error updating GoData case: %w", err)
		}
	}

	if err := goAPI.CreateCase(arg,
		godata.Options{URL: godataURL, Token: godataToken, OutbreakID: outbreakID}); err != nil {
		log.WithFields(log.Fields{
			"screening":    screening,
			"personalInfo": personalInfo,
			"caseID":       caseID,
		}).WithError(err).Error("failed to update godata case")
		return fmt.Errorf("error creating new GoData case: %w", err)
	}

	return nil
}
