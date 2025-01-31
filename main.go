// examples/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"quickauth/auth"
)

func main() {
	config := &auth.Config{
		ServiceAccountPath: os.Getenv("QS_GO_SA_JSON"),
		APIKey:            os.Getenv("QS_GO_API_KEY"),
	}
	
	if config.ServiceAccountPath == "" {
		log.Print("Please add env variable for QS_GO_SA_JSON as path to service account.json")
	}
	if config.APIKey == "" {
		log.Print("Please add env variable for QS_GO_API_KEY as api key of firebase project, find in Project settings in your firebase console")
	}

	client, err := auth.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	token, err := client.GenerateIDToken(context.Background(), "user@example.com")
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	fmt.Printf("Generated ID token: %s\n", token)
}
