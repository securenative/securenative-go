package models

import (
	. "github.com/securenative/securenative-go/securenative/enums"
	. "go/types"
)

type VerifyResult struct {
	RiskLevel RiskLevel
	Score int32
	Triggers Slice
}
