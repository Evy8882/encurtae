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
