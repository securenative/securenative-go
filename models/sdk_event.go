package models

import (
	"github.com/nu7hatch/gouuid"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/context"
	"github.com/securenative/securenative-go/errors"
	"github.com/securenative/securenative-go/utils"
	"time"
)

type SDKEvent struct {
	Context    *context.SecureNativeContext
	Rid        string
	EventType  string
	UserId     string
	UserTraits UserTraits
	Request    *RequestContext
	Timestamp  string
	Properties map[string]interface{}
}

func NewSDKEvent(eventOptions EventOptions, secureNativeOptions config.SecureNativeOptions) (SDKEvent, error) {
	if len(eventOptions.UserId) <= 0 || eventOptions.UserId == "" {
		return SDKEvent{}, &errors.SecureNativeInvalidOptionsError{Msg: "Invalid event structure; User Id is missing"}
	}

	if len(eventOptions.Event) <= 0 || eventOptions.Event == "" {
		return SDKEvent{}, &errors.SecureNativeInvalidOptionsError{Msg: "Invalid event structure; Event Type is missing"}
	}

	event := SDKEvent{}
	dateUtils := utils.DateUtils{}
	encryptionUtils := utils.EncryptionUtils{}

	if eventOptions.Context == nil {
		event.Context = &context.SecureNativeContext{}
	} else {
		event.Context = eventOptions.Context
	}

	clientToken := encryptionUtils.Decrypt(event.Context.ClientToken, secureNativeOptions.ApiKey)

	id, err := uuid.NewV4()
	if err != nil && id != nil {
		event.Rid = id.String()
	}

	event.EventType = eventOptions.Event
	event.UserId = eventOptions.UserId
	event.UserTraits = eventOptions.UserTraits

	event.Request = &RequestContext{
		Cid:      clientToken.Cid,
		Vid:      clientToken.Vid,
		Fp:       clientToken.Fp,
		Ip:       event.Context.Ip,
		RemoteIp: event.Context.RemoteIp,
		Headers:  event.Context.Headers,
		Url:      event.Context.Url,
		Method:   event.Context.Method,
	}

	t := time.Now().UTC()
	if eventOptions.Timestamp != nil {
		t = *eventOptions.Timestamp
	}

	event.Timestamp = dateUtils.ToTimestamp(t)
	event.Properties = eventOptions.Properties

	return event, nil
}
