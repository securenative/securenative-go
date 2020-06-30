package securenative

import (
	. "github.com/securenative/securenative-go/securenative/config"
	. "github.com/securenative/securenative-go/securenative/errors"
	. "github.com/securenative/securenative-go/securenative/http"
	. "github.com/securenative/securenative-go/securenative/models"
	"reflect"
)

type QueueItem struct {
	Url   string
	Body  string
	Retry bool
}

type EventManagerInterface interface {
	SendAsync(event SDKEvent, path string)
	SendSync(event SDKEvent, path string, retry bool) (map[string]string, error)
	StartEventPersist()
	StopEventPersist()
}

type EventManager struct {
	HttpClient   *HttpClient
	Queue        []QueueItem
	Options      SecureNativeOptions
	SendEnabled  bool
	Attempts     int16
	Coefficients []int32
	Interval     int32
}

func NewEventManager(options SecureNativeOptions, httpClient *HttpClient) (*EventManager, error) {
	if len(options.ApiKey) == 0 {
		return nil, &SecureNativeSDKError{Msg: "API key cannot be None, please get your API key from SecureNative console"}
	}

	var client *HttpClient
	if httpClient == nil {
		client = reflect.ValueOf(NewSecureNativeHttpClient(options)).Interface().(*HttpClient)
	} else {
		client = httpClient
	}

	coefficients := []int32{1, 1, 2, 3, 5, 8, 13}

	return &EventManager{
		HttpClient:   client,
		Queue:        []QueueItem{},
		Options:      options,
		SendEnabled:  false,
		Attempts:     0,
		Coefficients: coefficients,
		Interval:     options.Interval,
	}, nil
}

func (e *EventManager) SendAsync(event SDKEvent, path string) {
	// TODO implement me
}

func (e *EventManager) SendSync(event SDKEvent, path string, retry bool) (map[string]string, error) {
	// TODO implement me
	panic("implement me")
}

func (e *EventManager) StartEventPersist() {
	// TODO implement me
}

func (e *EventManager) StopEventPersist() {
	// TODO implement me
}