package tests

import (
	"github.com/jarcoal/httpmock"
	"github.com/securenative/securenative-go/events"
	"github.com/securenative/securenative-go/models"
	"testing"
	"time"
)

func getOptions() models.SDKEvent {
	return models.SDKEvent{
		Context:    nil,
		Rid:        "1",
		EventType:  "sn.user.login",
		UserId:     "1",
		UserTraits: models.UserTraits{Name: "Some User", Email: "email@securenative.com", Phone: "+12012673412"},
		Request:    &models.RequestContext{},
		Timestamp:  time.Now().Format("2006-01-02T15:04:05Z"),
		Properties: map[string]interface{}{},
	}
}

func TestShouldSuccessfullySendSyncEventWithStatusCode200(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	secureNativeOptions := getSecureNativeOptions()
	eventManager := events.NewEventManager(secureNativeOptions, nil)
	eventManager.StartEventPersist()
	defer eventManager.StopEventPersist()

	body := "{\"data\": true}"
	httpmock.RegisterResponder("POST", "https://api.securenative-stg.com/collector/api/v1/track", httpmock.NewStringResponder(200, body))

	data, _ := eventManager.SendSync(getOptions(), "track", false)

	if data["data"] != true {
		t.Errorf("Test Failed: data recived does not match: got: %s, expected: %t", data, true)
	}
}

func TestShouldSendSyncEventAndSailWhenStatusCode401(t *testing.T)  {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	secureNativeOptions := getSecureNativeOptions()
	eventManager := events.NewEventManager(secureNativeOptions, nil)
	eventManager.StartEventPersist()
	defer eventManager.StopEventPersist()

	httpmock.RegisterResponder("POST", "https://api.securenative-stg.com/collector/api/v1/track", httpmock.NewStringResponder(401, ""))

	_, err := eventManager.SendSync(getOptions(), "track", false)
	if err == nil {
		t.Errorf("Test Failed: expected to recieve 401 status code")
	}
}

func TestShouldSendSyncEventAndFailWhenStatusCode500(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	secureNativeOptions := getSecureNativeOptions()
	eventManager := events.NewEventManager(secureNativeOptions, nil)
	eventManager.StartEventPersist()
	defer eventManager.StopEventPersist()

	httpmock.RegisterResponder("POST", "https://api.securenative-stg.com/collector/api/v1/track", httpmock.NewStringResponder(500, ""))

	_, err := eventManager.SendSync(getOptions(), "track", false)
	if err == nil {
		t.Errorf("Test Failed: expected to recieve 500 status code")
	}
}