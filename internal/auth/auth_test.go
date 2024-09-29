package auth_test

import (
	"log"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	headers := map[string][]string{
		"Temp": []string{"123"},
	}

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		log.Fatalf("Oh no!")
	}

}
