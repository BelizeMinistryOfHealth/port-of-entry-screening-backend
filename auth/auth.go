package auth

import (
	"bz.moh.epi/poebackend/repository/firestore"
	"context"
	"fmt"
	"net/http"
	"strings"
)

// VerifyToken verifies that a JWT is valid. Returns error if the validation fails.
func VerifyToken(ctx context.Context, db firestore.DB, token string) error {
	_, err := db.AuthClient.VerifyIDToken(ctx, token)
	if err != nil {
		return fmt.Errorf("VerifyToken() failed: %w", err)
	}
	return nil
}

// JwtMiddleware is a middleware that verifies a JWT token
func JwtMiddleware(db firestore.DB, r *http.Request) error {
	ctx := r.Context()
	h := r.Header
	bearer := h.Get("Authorization")
	if len(strings.Trim(bearer, "")) == 0 {
		// No Authorization Token was provided
		//http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return fmt.Errorf("missing authorization header") //nolint:goerr113
	}
	bearerParts := strings.Split(bearer, " ")
	if bearerParts[0] != "Bearer" {
		// Wrong header format... return error
		//http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return fmt.Errorf("missing bearer header") //nolint:goerr113
	}
	token := bearerParts[1]
	return VerifyToken(ctx, db, token)
}
