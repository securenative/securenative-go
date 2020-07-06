package models

type VerifyResult struct {
	RiskLevel string
	Score     float64
	Triggers  []string
}
