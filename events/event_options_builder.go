package events

import (
	"fmt"
	"github.com/securenative/securenative-go/context"
	"github.com/securenative/securenative-go/errors"
	"github.com/securenative/securenative-go/models"
)

const MaxPropertiesSize = 10

type EventOptionsBuilder struct {
	EventOptions models.EventOptions
}

func NewEventOptionsBuilder(eventType string) *EventOptionsBuilder {
	options := models.EventOptions{
		Event:      eventType,
		UserId:     "",
		UserTraits: models.UserTraits{},
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

func (e *EventOptionsBuilder) WithUserTraits(userTraits models.UserTraits) *EventOptionsBuilder {
	e.EventOptions.UserTraits = userTraits
	return e
}

func (e *EventOptionsBuilder) WithContext(context *context.SecureNativeContext) *EventOptionsBuilder {
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

func (e *EventOptionsBuilder) Build() (models.EventOptions, error) {
	if len(e.EventOptions.Properties) > 0 && len(e.EventOptions.Properties) > MaxPropertiesSize {
		return models.EventOptions{}, &errors.SecureNativeInvalidOptionsError{Msg: fmt.Sprintf("You can have only up to %d custom properties", MaxPropertiesSize)}
	}

	return e.EventOptions, nil
}
