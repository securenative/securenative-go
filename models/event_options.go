package models

import (
	"github.com/securenative/securenative-go/context"
	"time"
)

type EventOptions struct {
	Event      string
	UserId     string
	UserTraits UserTraits
	Context    *context.SecureNativeContext
	Properties map[string]interface{}
	Timestamp  *time.Time
}
