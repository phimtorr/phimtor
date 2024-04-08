package auth

import (
	"os"

	"github.com/phimtorr/phimtor/server/admin/auth/ui"
)

func LoadFirebaseConfig() ui.FirebaseConfig {
	apiKey := os.Getenv("FIREBASE_API_KEY")
	if apiKey == "" {
		panic("FIREBASE_API_KEY is required")
	}
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	if projectID == "" {
		panic("FIREBASE_PROJECT_ID is required")
	}
	appID := os.Getenv("FIREBASE_APP_ID")
	if appID == "" {
		panic("FIREBASE_APP_ID is required")

	}
	measurementID := os.Getenv("FIREBASE_MEASUREMENT_ID")
	if measurementID == "" {
		panic("FIREBASE_MEASUREMENT_ID is required")
	}
	messagingSenderID := os.Getenv("FIREBASE_MESSAGING_SENDER_ID")
	if messagingSenderID == "" {
		panic("FIREBASE_MESSAGING_SENDER_ID is required")
	}

	return ui.FirebaseConfig{
		APIKey:            apiKey,
		ProjectID:         projectID,
		AppID:             appID,
		MeasurementID:     measurementID,
		MessagingSenderID: messagingSenderID,
	}
}
