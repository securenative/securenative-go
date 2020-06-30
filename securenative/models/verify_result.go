package models

type VerifyResult struct {
	RiskLevel string
	Score int32
	Triggers []string
}
