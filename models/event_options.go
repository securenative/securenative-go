package models

import (
	"github.com/securenative/securenative-go/context"
)

type EventOptions struct {
	Event      string
	UserId     string
	UserTraits UserTraits
	Context    *context.SecureNativeContext
	Properties map[string]string
	Timestamp  string
}
