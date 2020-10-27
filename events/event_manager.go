package events

import (
	"encoding/json"
	"fmt"
	"github.com/securenative/securenative-go"
	client "github.com/securenative/securenative-go/client"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/errors"
	"github.com/securenative/securenative-go/models"
	"github.com/securenative/securenative-go/utils"
	"io/ioutil"
	"net/http"
	"time"
)

var logger = securenative_go.GetLogger()

type QueueItem struct {
	Url   string
	Body  []byte
	Retry bool
}

type EventManagerInterface interface {
	SendAsync(event models.SDKEvent, path string)
	SendSync(event models.SDKEvent, path string, retry bool) (map[string]string, error)
	StartEventPersist()
	StopEventPersist()
}

type EventManager struct {
	HttpClient   *client.SecureNativeHttpClient
	Queue        []QueueItem
	Options      config.SecureNativeOptions
	SendEnabled  bool
	Attempt      int16
	Coefficients []int32
	Interval     int32
	Channel      chan struct{}
}

func NewEventManager(options config.SecureNativeOptions, httpClient *client.SecureNativeHttpClient) *EventManager {
	var c *client.SecureNativeHttpClient
	if httpClient == nil {
		c = client.NewSecureNativeHttpClient(options)
	} else {
		c = httpClient
	}

	coefficients := []int32{1, 1, 2, 3, 5, 8, 13}
	channel := make(chan struct{})

	return &EventManager{
		HttpClient:   c,
		Queue:        []QueueItem{},
		Options:      options,
		SendEnabled:  false,
		Attempt:      0,
		Coefficients: coefficients,
		Interval:     options.Interval,
		Channel:      channel,
	}
}

func (e *EventManager) SendAsync(event models.SDKEvent, path string) {
	if e.Options.Disable {
		logger.Warning("SDK is disabled. no operation will be performed")
		return
	}

	body, err := json.Marshal(e.serialize(event))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to marshal event body; %s", err))
		return
	}

	item := QueueItem{
		Url:   path,
		Body:  body,
		Retry: false,
	}
	e.Queue = append(e.Queue, item)
}

func (e *EventManager) SendSync(event models.SDKEvent, path string, retry bool) (map[string]interface{}, error) {
	if e.Options.Disable {
		logger.Warning("SDK is disabled. no operation will be performed")
		return nil, &errors.SecureNativeSDKIllegalStateError{Msg: "SDK is disabled. no operation will be performed"}
	}

	body, err := json.Marshal(e.serialize(event))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to marshal event body; %s", err))
		return nil, err
	}
	logger.Debug(fmt.Sprintf("Attempting to send event %s", body))

	res := e.HttpClient.Post(
		path,
		body,
	)

	if res.StatusCode != 200 {
		logger.Info(fmt.Sprintf("SecureNative failed to call endpoint %s with event %s. adding back to queue", path, event))
		item := QueueItem{
			Url:   path,
			Body:  body,
			Retry: retry,
		}
		e.Queue = append(e.Queue, item)
		return nil, fmt.Errorf("failed to send event; %d; %s", res.StatusCode, res.Status)
	}

	return readBody(res)
}

func (e *EventManager) StartEventPersist() {
	logger.Debug("Starting automatic event persistence")
	if e.Options.AutoSend || e.SendEnabled {
		e.SendEnabled = true
		go e.run()
	} else {
		logger.Debug("Automatic event persistence is disabled, you should persist events manually")
	}
}

func (e *EventManager) StopEventPersist() {
	if e.SendEnabled {
		logger.Debug("Attempting to stop automatic event persistence")
		e.flush()
		e.SendEnabled = false
		logger.Debug("Stopped event persistence")
	}
}

func (e *EventManager) flush() {
	for _, item := range e.Queue {
		e.HttpClient.Post(item.Url, item.Body)
	}
}

func (e *EventManager) run() {
	for true {
		if len(e.Queue) > 0 && e.SendEnabled {
			for i, item := range e.Queue {
				res := e.HttpClient.Post(item.Url, item.Body)
				e.Queue = removeItem(e.Queue, i)
				if res.StatusCode == 401 {
					item.Retry = false
				} else if res.StatusCode != 200 {
					item.Retry = true
				}

				logger.Debug(fmt.Sprintf("Event successfully sent; %s", item.Body))
				_, err := readBody(res)
				if err != nil {
					logger.Error(fmt.Sprintf("Failed to send event; %s", err))
					if item.Retry {
						if len(e.Coefficients) == int(e.Attempt+1) {
							e.Attempt = 0
						}

						backOff := e.Coefficients[e.Attempt] * e.Options.Interval
						logger.Debug(fmt.Sprintf("Automatic back-off of %d", backOff))
						e.SendEnabled = false
						time.Sleep(time.Duration(backOff))
						e.SendEnabled = true
					}
				}
			}
		}
		time.Sleep(time.Duration(e.Interval / 1000))
	}
}

func (e *EventManager) serialize(event models.SDKEvent) models.EventInput {
	dateUtils := utils.NewDateUtils()
	createdAt := dateUtils.ToTimestamp(time.Now())
	if len(event.UserTraits.CreatedAt) != 0 {
		createdAt = event.UserTraits.CreatedAt
	}

	serialized := models.EventInput{
		RequestID: event.Rid,
		EventType: event.EventType,
		UserID:    event.UserId,
		UserTraits: models.UserTraits{
			Name:      event.UserTraits.Name,
			Email:     event.UserTraits.Email,
			Phone:     event.UserTraits.Phone,
			CreatedAt: createdAt,
		},
		Request: models.RequestContext{
			Cid:         event.Request.Cid,
			Vid:         event.Request.Vid,
			Fp:          event.Request.Fp,
			Ip:          event.Request.Ip,
			RemoteIp:    event.Request.RemoteIp,
			Headers:     event.Request.Headers,
			Url:         event.Request.Url,
			Method:      event.Request.Method,
			ClientToken: event.Request.ClientToken,
		},
		Properties: event.Properties,
		Timestamp:  event.Timestamp,
	}
	return serialized
}

func readBody(response *http.Response) (map[string]interface{}, error) {
	var resBody map[string]interface{}

	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to read response body; %s", err))
		return nil, err
	}

	err = json.Unmarshal(b, &resBody)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to unmarshal response body; %s", err))
		return nil, err
	}

	return resBody, nil
}

func removeItem(s []QueueItem, index int) []QueueItem {
	return append(s[:index], s[index+1:]...)
}
