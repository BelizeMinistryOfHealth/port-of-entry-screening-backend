package firestore

import (
	fs "cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	auth "firebase.google.com/go/v4/auth"
	"fmt"
	"time"
)

// DB represents the database connection.
type DB struct {
	// The firestore client
	Client *fs.Client
	// The auth client. Used for verifying authentication
	AuthClient *auth.Client
	// FirebaseApp used for user administration
	FirebaseApp *firebase.App
	// Returns the current time. Defaults to time.Now().
	// Can be mocked for tests.
	Now       func() time.Time
	projectID string
}

// CreateFirestoreDB creates a new Firestore connection
func CreateFirestoreDB(ctx context.Context, projectID string) (*DB, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateFirestoreDB: could not create new firebase app: %w", err)
	}
	authClient, _ := app.Auth(ctx)

	c, err := fs.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("CreateFirestoreDB: could not create new firestore client: %w", err)
	}
	return &DB{
		Client:      c,
		AuthClient:  authClient,
		FirebaseApp: app,
		Now:         time.Now,
		projectID:   projectID,
	}, nil
}
