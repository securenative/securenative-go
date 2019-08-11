package models

type RiskResult struct {
	riskLevel string
	score     float64
	triggers  []string
}
