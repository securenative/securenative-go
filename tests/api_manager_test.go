package tests

import (
	"github.com/jarcoal/httpmock"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/context"
	"github.com/securenative/securenative-go/enums"
	"github.com/securenative/securenative-go/events"
	"github.com/securenative/securenative-go/models"
	"strconv"
	"testing"
)

func getContext() *context.SecureNativeContext {
	contextBuilder := context.NewSecureNativeContextBuilder()

	return contextBuilder.WithIp("127.0.0.1").
		WithClientToken("SECURED_CLIENT_TOKEN").
		WithHeaders(map[string]string{
			"user-agent": "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"}).Build()
}

func getEventOptions() models.EventOptions {
	eventOptionsBuilder := events.NewEventOptionsBuilder(enums.EventTypes.LogIn)
	options, _ := eventOptionsBuilder.
		WithUserId("USER_ID").
		WithUserTraits(models.UserTraits{
			Name:  "USER_NAME",
			Email: "USER_EMAIL",
		}).
		WithContext(getContext()).
		WithProperties(map[string]interface{}{"prop1": "CUSTOM_PARAM_VALUE", "prop2": true, "prop3": 3}).Build()

	return options
}

func getSecureNativeOptions() config.SecureNativeOptions {
	configBuilder := config.NewConfigurationBuilder()
	return configBuilder.
		WithApiKey("YOUR_API_KEY").
		WithAutoSend(true).
		WithInterval(10).
		WithApiUrl("https://api.securenative-stg.com/collector/api/v1").Build()
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

	apiManager.Track(getEventOptions())
	result := httpmock.GetCallCountInfo()

	if len(result) < 1 {
		t.Errorf("Test Failed: number of tracking post is: %d, expected: 1", len(result))
	}
}

func TestSecureNativeInvalidOptionsError(t *testing.T) {
	eventOptionsBuilder := events.NewEventOptionsBuilder(enums.EventTypes.LogIn)
	properties := map[string]interface{}{}
	for i := 1; i <= 12; i++ {
		properties[strconv.Itoa(i)] = i
	}
	_, err := eventOptionsBuilder.WithProperties(properties).Build()

	if err == nil {
		t.Error("Test Failed: expected SecureNativeInvalidOptionsError error to be thrown")
	}
}
