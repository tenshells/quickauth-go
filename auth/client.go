package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Client struct {
	config     *Config
	firebaseApp *firebase.App
}

func NewClient(config *Config) (*Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.ServiceAccountPath)
	
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return &Client{
		config:     config,
		firebaseApp: app,
	}, nil
}
