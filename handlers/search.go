package handlers

import (
	"bz.moh.epi/poebackend/auth"
	firesearch2 "bz.moh.epi/poebackend/repository/firesearch"
	"bz.moh.epi/poebackend/repository/firestore"
	"encoding/json"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type accessKeyResponse struct {
	AccessKey string `json:"accessKey"`
}

// AccessKeyHandler returns a firesearch access key
func AccessKeyHandler(db firestore.DB, w http.ResponseWriter, r *http.Request) {
	// get an idtoken.
	err := auth.JwtMiddleware(db, r)
	if err != nil {
		log.WithFields(log.Fields{
			"message": "JWT verification failed",
			"handler": "AccessKeyHandler",
		}).WithError(err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
	ctx := r.Context()
	log.Info("Creating Firesearch service...")
	firesearchService := firesearch2.CreateFiresearchService(
		"Persons Index",
		"persons_index",
		"PGIA")
	log.Info("retrieving access key")
	accessKeyService := firesearch.NewAccessKeyService(firesearchService.Client)
	keyReq := firesearch.GenerateKeyRequest{IndexPathPrefix: "firesearch/indexes/persons_index"}
	log.WithFields(log.Fields{
		"keyReq": keyReq,
	}).Info("Generating Access Key")
	keyResp, err := accessKeyService.GenerateKey(ctx, keyReq)
	if err != nil {
		log.WithFields(log.Fields{
			"message": "could not retrieve firesearch access key",
			"handler": "AccessKeyHandler",
		}).WithError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	accessKey := keyResp.AccessKey
	json.NewEncoder(w).Encode(accessKeyResponse{AccessKey: accessKey}) //nolint:errcheck,gosec

}

// ArrivalsStatAccessKeyHandler returns a firesearch access key
func ArrivalsStatAccessKeyHandler(db firestore.DB, w http.ResponseWriter, r *http.Request) {
	// get an idtoken.
	err := auth.JwtMiddleware(db, r)
	if err != nil {
		log.WithFields(log.Fields{
			"message": "JWT verification failed",
			"handler": "ArrivalsStatAccessKeyHandler",
		}).WithError(err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
	ctx := r.Context()
	log.Info("Creating Firesearch service...")
	firesearchService := firesearch2.CreateFiresearchService(
		"Arrivals Index",
		"arrivals_stat_index",
		"PGIA")
	log.Info("retrieving access key")
	accessKeyService := firesearch.NewAccessKeyService(firesearchService.Client)
	keyReq := firesearch.GenerateKeyRequest{IndexPathPrefix: "firesearch/indexes/arrivals_stat_index"}
	log.WithFields(log.Fields{
		"keyReq": keyReq,
	}).Info("Generating Access Key")
	keyResp, err := accessKeyService.GenerateKey(ctx, keyReq)
	if err != nil {
		log.WithFields(log.Fields{
			"message": "could not retrieve firesearch access key",
			"handler": "ArrivalsStatAccessKeyHandler",
		}).WithError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	accessKey := keyResp.AccessKey
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(accessKeyResponse{AccessKey: accessKey}) //nolint:errcheck,gosec

}
