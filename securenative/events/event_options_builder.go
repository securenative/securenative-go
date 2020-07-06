package events

import (
	"fmt"
	. "github.com/securenative/securenative-go/securenative/context"
	. "github.com/securenative/securenative-go/securenative/errors"
	. "github.com/securenative/securenative-go/securenative/models"
)

const MaxPropertiesSize = 10

type EventOptionsBuilder struct {
	EventOptions EventOptions
}

func NewEventOptionsBuilder(eventType string) *EventOptionsBuilder {
	options := EventOptions{
		Event:      eventType,
		UserId:     "",
		UserTraits: UserTraits{},
		Context:    nil,
		Properties: nil,
		Timestamp:  "",
	}

	return &EventOptionsBuilder{EventOptions: options}
}

func (e *EventOptionsBuilder) WithUserId(userId string) *EventOptionsBuilder {
	e.EventOptions.UserId = userId
	return e
}

func (e *EventOptionsBuilder) WithUserTraits(userTraits UserTraits) *EventOptionsBuilder {
	e.EventOptions.UserTraits = userTraits
	return e
}

func (e *EventOptionsBuilder) WithContext(context *SecureNativeContext) *EventOptionsBuilder {
	e.EventOptions.Context = context
	return e
}

func (e *EventOptionsBuilder) WithProperties(properties map[string]string) *EventOptionsBuilder {
	e.EventOptions.Properties = properties
	return e
}

func (e *EventOptionsBuilder) WithTimestamp(timestamp string) *EventOptionsBuilder {
	e.EventOptions.Timestamp = timestamp
	return e
}

func (e *EventOptionsBuilder) Build() (EventOptions, error) {
	if len(e.EventOptions.Properties) > 0 && len(e.EventOptions.Properties) > MaxPropertiesSize {
		return EventOptions{}, &SecureNativeInvalidOptionsError{Msg: fmt.Sprintf("You can have only up to %d custom properties", MaxPropertiesSize)}
	}

	return e.EventOptions, nil
}
