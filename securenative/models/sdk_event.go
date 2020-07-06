package models

import (
	"github.com/nu7hatch/gouuid"
	"github.com/securenative/securenative-go/securenative/config"
	"github.com/securenative/securenative-go/securenative/context"
	"github.com/securenative/securenative-go/securenative/utils"
)

type SDKEvent struct {
	Context    *context.SecureNativeContext
	Rid        string
	EventType  string
	UserId     string
	UserTraits UserTraits
	Request    *RequestContext
	Timestamp  string
	Properties map[string]string
}

func NewSDKEvent(eventOptions EventOptions, secureNativeOptions config.SecureNativeOptions) SDKEvent {
	event := SDKEvent{}
	dateUtils := utils.DateUtils{}
	encryptionUtils := utils.EncryptionUtils{}
	contextBuilder := context.NewSecureNativeContextBuilder()

	if eventOptions.Context == nil {
		event.Context = eventOptions.Context
	} else {
		event.Context = contextBuilder.Build()
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

	event.Timestamp = dateUtils.ToTimestamp(eventOptions.Timestamp)
	event.Properties = eventOptions.Properties

	return event
}
