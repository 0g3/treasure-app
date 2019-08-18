package firebaseutil

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewAuthClient(credentialFilePath string) (*auth.Client, error) {
	opt := option.WithCredentialsFile(credentialFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app.Auth(context.Background())
}
