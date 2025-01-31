package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const firebaseAuthURL = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken"

func (c *Client) GenerateIDToken(ctx context.Context, email string) (string, error) {
	// Get Auth client
	authClient, err := c.firebaseApp.Auth(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting Auth client: %v", err)
	}

	// Generate custom token first
	customToken, err := authClient.CustomToken(ctx, email)
	if err != nil {
		return "", fmt.Errorf("error generating custom token: %v", err)
	}

	// Exchange custom token for ID token
	return c.exchangeCustomToken(customToken)
}

func (c *Client) exchangeCustomToken(customToken string) (string, error) {
	reqBody := SignInRequest{
		Token:             customToken,
		ReturnSecureToken: true,
	}
	
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	url := fmt.Sprintf("%s?key=%s", firebaseAuthURL, c.config.APIKey)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error response: %s", string(body))
	}

	var signInResp SignInResponse
	if err := json.Unmarshal(body, &signInResp); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	return signInResp.IDToken, nil
}
