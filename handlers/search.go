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
	err := auth.JwtMiddleware(db, r)
	if err != nil {
		log.WithFields(log.Fields{
			"message": "JWT verification failed",
			"handler": "AccessKeyHandler",
		}).WithError(err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
	ctx := r.Context()
	firesearchService := firesearch2.CreateFiresearchService(
		"Persons Index",
		"persons_index",
		"PGIA")
	accessKeyService := firesearch.NewAccessKeyService(firesearchService.Client)
	keyReq := firesearch.GenerateKeyRequest{IndexPathPrefix: "firesearch/indexes/persons_index"}
	keyResp, err := accessKeyService.GenerateKey(ctx, keyReq)
	if err != nil {
		log.WithFields(log.Fields{
			"message": "could not retrieve firesearch access key",
			"handler": "AccessKeyHandler",
		}).WithError(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	accessKey := keyResp.AccessKey
	json.NewEncoder(w).Encode(accessKeyResponse{AccessKey: accessKey})

	//results, searchErr := personStore.SearchByName(ctx, accessKey, "", "")
	//if searchErr != nil {
	//	log.WithFields(log.Fields{
	//		"handler": "AccessKeyHandler",
	//		"message": "searching for person by name failed",
	//	}).WithError(searchErr)
	//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	//}
	//
	//if jsonErr := json.NewEncoder(w).Encode(results); jsonErr != nil {
	//	log.WithFields(log.Fields{
	//		"handler": "AccessKeyHandler",
	//		"message": "encoding search results failed",
	//	}).WithError(jsonErr)
	//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	//}

}
