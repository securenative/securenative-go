package tests

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/enums"
	"github.com/securenative/securenative-go/models"
	"testing"
)

func TestEventCreatedWithoutUserIdThrow(t *testing.T) {
	eventOptions := models.EventOptions{Event: enums.EventTypes.LogIn}
	options := config.SecureNativeOptions{}

	_, err := models.NewSDKEvent(eventOptions, options)

	if err == nil {
		t.Error("Test Failed: expected SecureNativeInvalidOptionsError error to be thrown")
	}
}

func TestEventCreatedWithoutEventTypeThrow(t *testing.T) {
	eventOptions := models.EventOptions{UserId: "1234"}
	options := config.SecureNativeOptions{}

	_, err := models.NewSDKEvent(eventOptions, options)

	if err == nil {
		t.Error("Test Failed: expected SecureNativeInvalidOptionsError error to be thrown")
	}
}