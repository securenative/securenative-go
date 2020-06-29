package models

import (
	. "github.com/securenative/securenative-go/securenative/context"
	. "github.com/securenative/securenative-go/securenative/enums"
	. "go/types"
)

type EventOptions struct {
	Event      EventTypes
	UserId     string
	UserTraits *UserTraits
	Context    *SecureNativeContext
	Properties Slice
	Timestamp  string
}
