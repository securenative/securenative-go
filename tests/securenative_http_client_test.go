package tests

import (
	"github.com/jarcoal/httpmock"
	"github.com/securenative/securenative-go/securenative/client"
	"testing"
)

func TestShouldMakeSimplePostCall(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	secureNativeOptions := getSecureNativeOptions()
	httpClient := client.SecureNativeHttpClient{Options: secureNativeOptions}

	expected := "{\"event\": \"SOME_EVENT_NAME\"}"
	httpmock.RegisterResponder("POST", "https://api.securenative-stg.com/collector/api/v1/track", httpmock.NewStringResponder(200, expected))

	result := httpClient.Post("track", []byte(expected))

	if result.StatusCode != 200 && result.Status != "200 OK" {
		t.Errorf("Test Failed: data recived does not match: got: %d, expected: %d", result.StatusCode, 200)
	}
}
