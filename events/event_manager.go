package events

import (
	"encoding/json"
	"fmt"
	"github.com/securenative/securenative-go/client"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/errors"
	"github.com/securenative/securenative-go/logger"
	"github.com/securenative/securenative-go/models"
	"github.com/securenative/securenative-go/utils"
	"io/ioutil"
	"net/http"
	"time"
)

var log = logger.GetLogger()

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
		log.Warning("SDK is disabled. no operation will be performed")
		return
	}

	body, err := json.Marshal(e.serialize(event))
	if err != nil {
		log.Error(fmt.Sprintf("Failed to marshal event body; %s", err))
		return
	}

	item := QueueItem{
		Url:   path,
		Body:  body,
		Retry: false,
	}
	e.Queue = append(e.Queue, item)
}

func (e *EventManager) SendSync(event models.SDKEvent, path string) (map[string]interface{}, error) {
	if e.Options.Disable {
		log.Warning("SDK is disabled. no operation will be performed")
		return nil, &errors.SecureNativeSDKIllegalStateError{Msg: "SDK is disabled. no operation will be performed"}
	}

	body, err := json.Marshal(e.serialize(event))
	if err != nil {
		log.Error(fmt.Sprintf("Failed to marshal event body; %s", err))
		return nil, err
	}
	log.Debug(fmt.Sprintf("Attempting to send event %s", body))

	res, err := e.HttpClient.Post(
		path,
		body,
	)

	if err != nil || res != nil && res.StatusCode != 200 || res == nil {
		log.Info(fmt.Sprintf("SecureNative failed to call endpoint %s with event %s. adding back to queue", path, event))
		return nil, fmt.Errorf("failed to send event; %s", err)
	}

	return readBody(res)
}

func (e *EventManager) StartEventPersist() {
	log.Debug("Starting automatic event persistence")
	if e.Options.AutoSend || e.SendEnabled {
		e.SendEnabled = true
		go e.run()
	} else {
		log.Debug("Automatic event persistence is disabled, you should persist events manually")
	}
}

func (e *EventManager) StopEventPersist() {
	if e.SendEnabled {
		log.Debug("Attempting to stop automatic event persistence")
		e.flush()
		e.SendEnabled = false
		log.Debug("Stopped event persistence")
	}
}

func (e *EventManager) flush() {
	for _, item := range e.Queue {
		_, _ = e.HttpClient.Post(item.Url, item.Body)
	}
}

func (e *EventManager) run() {
	for true {
		if len(e.Queue) > 0 && e.SendEnabled {
			for i, item := range e.Queue {
				res, err := e.HttpClient.Post(item.Url, item.Body)
				e.Queue = removeItem(e.Queue, i)

				if err != nil || res != nil && res.StatusCode != 200{
					item.Retry = true
					log.Error(fmt.Sprintf("Failed to send event; %s", err))
					e.backOffSend(item)
				} else if res != nil && res.StatusCode == 401 {
					item.Retry = false
					log.Error(fmt.Sprintf("Failed to send event; %s", err))
				} else {
					_, err = readBody(res)
					if err != nil {
						log.Error(fmt.Sprintf("Failed to send event; %s", err))
						e.backOffSend(item)
					}
					log.Debug(fmt.Sprintf("Event successfully sent; %s", item.Body))
				}
			}
		}
		time.Sleep(time.Duration(e.Interval / 1000))
	}
}

func (e *EventManager) backOffSend(item QueueItem) {
	if item.Retry {
		if len(e.Coefficients) == int(e.Attempt+1) {
			e.Attempt = 0
		}

		backOff := e.Coefficients[e.Attempt] * e.Options.Interval
		log.Debug(fmt.Sprintf("Automatic back-off of %d", backOff))
		e.SendEnabled = false
		time.Sleep(time.Duration(backOff))
		e.SendEnabled = true
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
		log.Error(fmt.Sprintf("Failed to read response body; %s", err))
		return nil, err
	}

	err = json.Unmarshal(b, &resBody)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to unmarshal response body; %s", err))
		return nil, err
	}

	return resBody, nil
}

func removeItem(s []QueueItem, index int) []QueueItem {
	return append(s[:index], s[index+1:]...)
}
