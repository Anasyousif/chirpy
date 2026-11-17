package auth

import (
	"testing"
	"time"
    "net/http"
	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "super-secret-key"
	expiresIn := time.Hour

	// 1. Create token
	tokenStr, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	// 2. Validate token
	parsedID, err := ValidateJWT(tokenStr, secret)
	if err != nil {
		t.Fatalf("unexpected error validating JWT: %v", err)
	}

	if parsedID != userID {
		t.Errorf("expected user ID %v, got %v", userID, parsedID)
	}
}

func TestValidateJWTWrongSecret(t *testing.T) {
	userID := uuid.New()
	secret := "super-secret-key"
	wrongSecret := "wrong-secret-key"
	expiresIn := time.Hour

	tokenStr, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	// Validate with wrong secret should fail
	_, err = ValidateJWT(tokenStr, wrongSecret)
	if err == nil {
		t.Error("expected error when validating with wrong secret, got nil")
	}
}

func TestValidateJWTExpired(t *testing.T) {
	userID := uuid.New()
	secret := "super-secret-key"
	// Set expiration to negative so it's already expired
	expiresIn := -time.Hour

	tokenStr, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	// Validate expired token should fail
	_, err = ValidateJWT(tokenStr, secret)
	if err == nil {
		t.Error("expected error when validating expired token, got nil")
	}
}
func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-secret-token-123")

	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token != "my-secret-token-123" {
		t.Errorf("expected 'my-secret-token-123', got '%s'", token)
	}
}