package auth

import (
	"crypto/subtle"
	"errors"
	"time"

	"github.com/google/uuid"
)

type APIKey struct {
	Key       string
	ExpiresAt time.Time
}

var validAPIKeys map[string]APIKey

// Function to create a new APIKey
func init() {
	validAPIKeys = make(map[string]APIKey)
	AddNewAPIKey(30 * 24 * time.Hour)
}

func AddNewAPIKey(duration time.Duration) string {
	newKey := uuid.New().String()
	validAPIKeys[newKey] = APIKey{
		Key:       newKey,
		ExpiresAt: time.Now().Add(duration),
	}
	return newKey
}

func ValidateAPIKey(apiKey string) error {
	if apiKey == "" {
		return errors.New("API key is required")
	}

	key, exists := validAPIKeys[apiKey]
	if !exists || time.Now().After(key.ExpiresAt) {
		return errors.New("invalid or expired API key")
	}

	// Use constant-time comparison for safety, prevents brute force cracking for the API Key
	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(key.Key)) != 1 {
		return errors.New("invalid API key")
	}

	return nil
}
