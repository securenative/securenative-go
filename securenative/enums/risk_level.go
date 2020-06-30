package enums

type RiskLevelEnum struct {
	Low    string
	Medium string
	High   string
}

var RiskLevel = RiskLevelEnum{
	Low:    "low",
	Medium: "medium",
	High:   "high",
}
