package tests

import (
	"github.com/jarcoal/httpmock"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/context"
	"github.com/securenative/securenative-go/enums"
	"github.com/securenative/securenative-go/events"
	"github.com/securenative/securenative-go/models"
	"testing"
)

func getContext() *context.SecureNativeContext {
	return &context.SecureNativeContext{
		ClientToken: "SECURED_CLIENT_TOKEN",
		Ip:          "127.0.0.1",
		Headers: map[string]string{
			"user-agent": "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"},
	}
}

func getEventOptions() models.EventOptions {
	return models.EventOptions{
		UserId: "USER_ID",
		Event: enums.EventTypes.LogIn,
		UserTraits: models.UserTraits{
			Name:  "USER_NAME",
			Email: "USER_EMAIL",
			Phone: "+12012673412",
		},
		Context:    getContext(),
		Properties: map[string]interface{}{"prop1": "CUSTOM_PARAM_VALUE", "prop2": true, "prop3": 3},
	}
}

func getSecureNativeOptions() config.SecureNativeOptions {
	options := config.DefaultSecureNativeOptions()
	options.ApiKey = "YOUR_API_KEY"
	options.AutoSend = true
	options.Interval = 10
	options.ApiUrl = "https://api.securenative-stg.com/collector/api/v1"
	return options
}

func TestTrackEvent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	secureNativeOptions := getSecureNativeOptions()
	eventManager := events.NewEventManager(secureNativeOptions, nil)
	eventManager.StartEventPersist()
	defer eventManager.StopEventPersist()
	apiManager := events.NewApiManager(eventManager, secureNativeOptions)

	expected := "{\"eventType\":\"sn.user.login\",\"userId\":\"USER_ID\",\"userTraits\":{\"name\":\"USER_NAME\",\"email\":\"USER_EMAIL\",\"createdAt\":nil},\"request\":{\"cid\":nil,\"vid\":nil,\"fp\":nil,\"ip\":\"127.0.0.1\",\"remoteIp\":nil,\"headers\":{\"user-agent\":\"Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405\"},\"url\":nil,\"method\":nil},\"properties\":{\"prop2\":true,\"prop1\":\"CUSTOM_PARAM_VALUE\",\"prop3\":3}}"
	httpmock.RegisterResponder("POST", "https://api.securenative-stg.com/collector/api/v1/track", httpmock.NewStringResponder(200, expected))

	_ = apiManager.Track(getEventOptions())
	result := httpmock.GetCallCountInfo()

	if len(result) < 1 {
		t.Errorf("Test Failed: number of tracking post is: %d, expected: 1", len(result))
	}
}

func TestTimeoutOnPost(t *testing.T) {
	secureNativeOptions := getSecureNativeOptions()
	secureNativeOptions.Timeout = 1

	eventManager := events.NewEventManager(secureNativeOptions, nil)
	eventManager.StartEventPersist()
	defer eventManager.StopEventPersist()
	apiManager := events.NewApiManager(eventManager, secureNativeOptions)

	httpmock.RegisterResponder("POST", "https://api.securenative-stg.com/collector/api/v1/verify", httpmock.NewStringResponder(408, ""))

	verifyResult, _ := apiManager.Verify(getEventOptions())
	expected := models.VerifyResult{
		RiskLevel: enums.RiskLevel.Low,
		Score: 0,
		Triggers: []string{},
	}

	if len(verifyResult.Triggers) != len(expected.Triggers) {
		t.Errorf("Test Failed: length of triggers is: %d, expected: %d", len(verifyResult.Triggers), len(expected.Triggers))
	}

	if verifyResult.Score != expected.Score {
		t.Errorf("Test Failed: score is: %v, expected: %v", verifyResult.Score, expected.Score)
	}

	if verifyResult.RiskLevel != expected.RiskLevel {
		t.Errorf("Test Failed: risk level is: %v, expected: %v", verifyResult.RiskLevel, expected.RiskLevel)
	}
}
