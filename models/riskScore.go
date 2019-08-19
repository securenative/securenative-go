package models

type RiskResult struct {
	RiskLevel string
	Score     float64
	Triggers  []string
}
