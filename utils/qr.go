package utils

import (
	"fmt"
	"os"
)

func GenerateBoxQR(boxID string) string {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173/api/v1"
	}
	targetURL := fmt.Sprintf("%s/boxes/%s", frontendURL, boxID)

	return targetURL
}
