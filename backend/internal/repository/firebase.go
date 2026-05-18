package repository

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	App  *firebase.App
	Auth *auth.Client
}

func NewFirebaseService() (*FirebaseService, error) {

	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase-admin-sdk.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return &FirebaseService{
		App:  app,
		Auth: authClient,
	}, nil
}

func (s *FirebaseService) DeleteUrl(ctx context.Context, id string) error {
	if s == nil || s.App == nil {
		return fmt.Errorf("firebase app not initialized")
	}

	client, err := s.App.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error getting Firestore client: %v", err)
	}
	defer client.Close()

	_, err = client.Collection("urls").Doc(id).Delete(ctx)
	if err != nil {
		return fmt.Errorf("error deleting URL: %v", err)
	}

	return nil
}
