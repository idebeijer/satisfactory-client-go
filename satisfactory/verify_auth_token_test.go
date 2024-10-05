package satisfactory

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_VerifyAuthenticationToken_Success(t *testing.T) {
	// Mock server to simulate a valid token response
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Verify that the Authorization header is present and valid
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer valid-token" {
			t.Errorf("Expected Authorization header 'Bearer valid-token', got '%s'", authHeader)
		}

		// Respond with a successful CommonResponse
		respData := CommonResponse{
			Data: json.RawMessage(`{}`), // Empty data for success
		}
		respBytes, _ := json.Marshal(respData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, nil, true)
	client.SetAuthToken("valid-token")

	resp, err := client.VerifyAuthenticationToken(context.Background())
	if err != nil {
		t.Fatalf("VerifyAuthenticationToken returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClient_VerifyAuthenticationToken_InvalidToken(t *testing.T) {
	// Mock server to simulate an invalid token response
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Verify that the Authorization header is present
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer invalid-token" {
			t.Errorf("Expected Authorization header 'Bearer invalid-token', got '%s'", authHeader)
		}

		// Respond with an error CommonResponse
		respData := CommonResponse{
			ErrorCode:    "invalid_token",
			ErrorMessage: "The provided authentication token is invalid.",
		}
		respBytes, _ := json.Marshal(respData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(respBytes)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewClient(server.URL, nil, true)
	client.SetAuthToken("invalid-token")

	resp, err := client.VerifyAuthenticationToken(context.Background())
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	apiErr, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("Expected ErrorResponse, got %T", err)
	}

	if apiErr.ErrorCode != "invalid_token" {
		t.Errorf("Expected ErrorCode 'invalid_token', got '%s'", apiErr.ErrorCode)
	}

	if apiErr.ErrorMessage != "The provided authentication token is invalid." {
		t.Errorf("Expected ErrorMessage 'The provided authentication token is invalid.', got '%s'", apiErr.ErrorMessage)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}
}

func TestClient_VerifyAuthenticationToken_Integration_Success(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode.")
	}

	client := NewClient("https://localhost:7777", nil, true)
	client.SetAuthToken("valid-token") // TODO: Replace with a valid token, through login or other means

	resp, err := client.VerifyAuthenticationToken(context.Background())
	if err != nil {
		t.Fatalf("VerifyAuthenticationToken returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClient_VerifyAuthenticationToken_Integration_InvalidToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode.")
	}

	client := NewClient("https://localhost:7777", nil, true)
	client.SetAuthToken("invalid-token")

	resp, err := client.VerifyAuthenticationToken(context.Background())
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	apiErr, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("Expected ErrorResponse, got %T", err)
	}

	// Adjust the expected error code and message according to your API's response
	if apiErr.ErrorCode != "invalid_token" {
		t.Errorf("Expected ErrorCode 'invalid_token', got '%s'", apiErr.ErrorCode)
	}

	if apiErr.ErrorMessage == "" {
		t.Errorf("Expected an error message, got empty string")
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}
}
